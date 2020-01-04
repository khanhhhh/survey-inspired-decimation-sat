package sat

import "math"

type propagationGraph struct {
	piMap  map[edge][3]message // variable -> clause
	etaMap map[edge]message    // clause -> variable
}

func (ins *instance) makePropagationGraph() *propagationGraph {
	out := &propagationGraph{
		make(map[edge][3]message),
		make(map[edge]message),
	}
	for _, e := range ins.allEdges() {
		out.piMap[e] = [3]message{0.5, 0.5, 0.5}
		out.etaMap[e] = 0.5
	}
	return out
}

func (ins *instance) iteratePropagationGraph(g *propagationGraph, smooth float64, tolerance float64) (bool, *propagationGraph) {
	var converged bool = true
	out := &propagationGraph{
		make(map[edge][3]message),
		make(map[edge]message),
	}
	for _, e := range ins.allEdges() {
		i := e.variable
		a := e.clause
		// eta
		var eta message = 1
		for j := range ins.clauseMap[a] {
			if j != i {
				triplet := g.piMap[edge{j, a}]
				eta *= triplet[0] / (triplet[0] + triplet[1] + triplet[2])
			}
		}
		if abs(eta-g.etaMap[e]) > tolerance {
			converged = false
		}
		// detect nan
		if math.IsNaN(eta) {
			return true, g
		}
		out.etaMap[e] = eta
		// pi
		var productAgree message = 1
		var productDisagree message = 1
		for _, b := range ins.clauseAgree(e) {
			productAgree *= 1 - g.etaMap[edge{i, b}]
		}
		for _, b := range ins.clauseDisagree(e) {
			productDisagree *= 1 - g.etaMap[edge{i, b}]
		}
		var triplet [3]message
		triplet[0] = (1 - smooth*productDisagree) * productAgree
		triplet[1] = (1 - smooth*productAgree) * productDisagree
		triplet[2] = smooth * productAgree * productDisagree
		out.piMap[e] = triplet
		// detect nan
		if math.IsNaN(triplet[0]) || math.IsNaN(triplet[1]) || math.IsNaN(triplet[2]) {
			return true, g
		}
	}
	return converged, out
}
