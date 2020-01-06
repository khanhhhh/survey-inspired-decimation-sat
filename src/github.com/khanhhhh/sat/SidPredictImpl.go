package sat

import (
	"math"
)

var tolerance float64 = 0.001
var smooth float64 = 1.0
var iterMul float64 = 100

func (ins *instance) SidPredict() (converged bool, nonTrivialCover bool, variable variable, value bool) {
	// survey propagation
	var absoluteEtaChange float64 = 1
	graph := ins.makeSurveyPropagationGraph()
	{
		numIterations := 1 + int(iterMul*math.Log2(float64(ins.capVariables())))
		for iteration := 0; iteration < numIterations; iteration++ {
			absoluteEtaChange, graph = ins.iterateSurveyPropagationGraph(graph, smooth)
			if absoluteEtaChange < tolerance {
				break
			}
		}
	}
	// converge ?
	{
		if absoluteEtaChange > tolerance {
			converged = false
		} else {
			converged = true
			nonTrivialCover, variable, value = ins.surveyDecimation(graph, smooth)
		}
	}
	return converged, nonTrivialCover, variable, value
}
