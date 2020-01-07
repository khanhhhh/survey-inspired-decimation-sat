package surveydecimation

import "github.com/khanhhhh/sat/guesser/surveydecimation/rational"

import "github.com/khanhhhh/sat/instance"

type surveyPropagationGraph struct {
	piMap  map[edge][3]rational.Rat // variable -> clause
	etaMap map[edge]rational.Rat    // clause -> variable
}

// makeSurveyPropagationGraph :
// Make an empty Survey Propagation Graph
func makeSurveyPropagationGraph(ins instance.Instance) (graph *surveyPropagationGraph) {
	graph = &surveyPropagationGraph{
		make(map[edge][3]rational.Rat),
		make(map[edge]rational.Rat),
	}
	for _, edge := range allEdges(ins) {
		graph.piMap[edge] = [3]rational.Rat{
			rational.FromInt(1, 2),
			rational.FromInt(1, 2),
			rational.FromInt(1, 2),
		}
		graph.etaMap[edge] = rational.FromInt(1, 2)
	}
	return graph
}

// iterateSurveyPropagationGraph :
// Iterate clauseA Survey Propagation Graph
func iterateSurveyPropagationGraph(ins instance.Instance, graphIn *surveyPropagationGraph, smooth float64) (absoluteEtaChange float64, graphOut *surveyPropagationGraph) {
	// initialize etaChange to 0
	absoluteEtaChange = 0
	// make empty graphOut
	graphOut = &surveyPropagationGraph{
		make(map[edge][3]rational.Rat),
		make(map[edge]rational.Rat),
	}
	// calculate graphOut value for all edges
	for _, edge := range allEdges(ins) {
		variableI := edge.variable
		clauseA := edge.clause
		// eta
		{
			var eta = rational.FromInt(1, 1)
			for variableJ := range ins.ClauseMap()[clauseA] {
				if variableJ != variableI {
					triplet := graphIn.piMap[newEdge(variableJ, clauseA)]
					sum := rational.Add(rational.Add(triplet[0], triplet[1]), triplet[2])
					eta *= rational.Div(triplet[0], sum)
				}
			}
			// detect nan : if sum triplet == 0 => eta = NaN
			//if math.IsNaN(eta) {
			//	panic("eta: NaN")
			//}
			if rational.Abs(rational.Sub(eta, graphIn.etaMap[edge])).ToFloat() > absoluteEtaChange {
				absoluteEtaChange = rational.Abs(rational.Sub(eta, graphIn.etaMap[edge])).ToFloat()
			}
			graphOut.etaMap[edge] = eta
		}
		// pi
		{
			var productAgree = rational.FromInt(1, 1)
			var productDisagree = rational.FromInt(1, 1)
			for _, clauseB := range clauseAgree(ins, edge) {
				productAgree = rational.Mul(
					productAgree,
					rational.Sub(rational.FromInt(1, 1), graphIn.etaMap[newEdge(variableI, clauseB)]),
				)
			}
			for _, clauseB := range clauseDisagree(ins, edge) {
				productDisagree = rational.Mul(
					productDisagree,
					rational.Sub(rational.FromInt(1, 1), graphIn.etaMap[newEdge(variableI, clauseB)]),
				)
			}
			var triplet [3]rational.Rat
			smoothConst := rational.FromFloat(smooth)
			triplet[0] = rational.Mul(
				productAgree,
				rational.Sub(rational.FromInt(1, 1), rational.Mul(smoothConst, productDisagree)),
			)
			triplet[1] = rational.Mul(
				productDisagree,
				rational.Sub(rational.FromInt(1, 1), rational.Mul(smoothConst, productAgree)),
			)
			triplet[2] = rational.Mul(
				rational.FromFloat(smooth),
				rational.Mul(productAgree, productDisagree),
			)
			// detect nan
			//if math.IsNaN(triplet[0]) || math.IsNaN(triplet[1]) || math.IsNaN(triplet[2]) {
			//	panic("triplet: NaN")
			//}
			// detect zero
			if triplet[0]+triplet[1]+triplet[2] == 0 {
				panic("triplet: Zero")
			}
			graphOut.piMap[edge] = triplet
		}
	}
	return absoluteEtaChange, graphOut
}
