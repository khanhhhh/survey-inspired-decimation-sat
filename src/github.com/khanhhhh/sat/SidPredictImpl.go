package sat

import (
	"math"
)

var tolerance float64 = 0.001
var smooth float64 = 1.0
var iterMul float64 = 100

// SidPredict :
// Survey Inspired Decimation: Predict the value of a variable
func (ins *instance) SidPredict() (converged bool, nonTrivialCover bool, variableOut variable, valueOut bool) {
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
	// converged ?
	{
		if absoluteEtaChange > tolerance {
			converged = false
		} else {
			converged = true
			nonTrivialCover, variableOut, valueOut = ins.surveyDecimation(graph, smooth)
		}
	}
	return converged, nonTrivialCover, variableOut, valueOut
}
