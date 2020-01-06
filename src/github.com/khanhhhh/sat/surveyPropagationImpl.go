package sat

import "math"

type surveyPropagationGraph struct {
	piMap  map[edge][3]message // variable -> clause
	etaMap map[edge]message    // clause -> variable
}

// Make the empty Survey Propagation Graph
func (ins *instance) makeSurveyPropagationGraph() (graph *surveyPropagationGraph) {
	graph = &surveyPropagationGraph{
		make(map[edge][3]message),
		make(map[edge]message),
	}
	for _, e := range ins.allEdges() {
		graph.piMap[e] = [3]message{0.5, 0.5, 0.5}
		graph.etaMap[e] = 0.5
	}
	return graph
}

// Iterate Survey Propagation Graph
func (ins *instance) iterateSurveyPropagationGraph(graphIn *surveyPropagationGraph, smooth float64) (absoluteEtaChange message, graphOut *surveyPropagationGraph) {
	// initialize etaChange to 0
	absoluteEtaChange = 0
	// make empty graphOut
	graphOut = &surveyPropagationGraph{
		make(map[edge][3]message),
		make(map[edge]message),
	}
	// calculate graphOut value for all edges
	for _, e := range ins.allEdges() {
		i := e.variable
		a := e.clause
		// eta
		{
			var eta message = 1
			for j := range ins.clauseMap[a] {
				if j != i {
					triplet := graphIn.piMap[edge{j, a}]
					eta *= triplet[0] / (triplet[0] + triplet[1] + triplet[2])
				}
			}
			if abs(eta-graphIn.etaMap[e]) > absoluteEtaChange {
				absoluteEtaChange = abs(eta - graphIn.etaMap[e])
			}
			// detect nan
			if math.IsNaN(eta) {
				panic("NaN")
			}
			graphOut.etaMap[e] = eta
		}
		// pi
		{
			var productAgree message = 1
			var productDisagree message = 1
			for _, b := range ins.clauseAgree(e) {
				productAgree *= 1 - graphIn.etaMap[edge{i, b}]
			}
			for _, b := range ins.clauseDisagree(e) {
				productDisagree *= 1 - graphIn.etaMap[edge{i, b}]
			}
			var triplet [3]message
			triplet[0] = (1 - smooth*productDisagree) * productAgree
			triplet[1] = (1 - smooth*productAgree) * productDisagree
			triplet[2] = smooth * productAgree * productDisagree
			graphOut.piMap[e] = triplet
			// detect nan
			if math.IsNaN(triplet[0]) || math.IsNaN(triplet[1]) || math.IsNaN(triplet[2]) {
				panic("NaN")
			}
		}
	}
	return absoluteEtaChange, graphOut
}
