// Elitza Maneva, Elchanan Mossel, and Martin J. Wainwright. 2007. A new look at survey propagation and its generalizations. J. ACM 54, 4 (July 2007), 17–es. DOI:https://doi.org/10.1145/1255443.1255445

package surveydecimation

import (
	"math"

	"github.com/khanhhhh/sat/instance"
)

var tolerance float64 = 0.001
var iterMul float64 = 100

// Guess :
// Survey Inspired Decimation: Predict the value of a variable
func Guess(ins instance.Instance, smooth float64) (converged bool, nonTrivialCover bool, variableOut instance.Variable, valueOut bool) {
	// survey propagation
	var absoluteEtaChange float64 = 1
	graph := makeSurveyPropagationGraph(ins)
	{
		numIterations := 1 + int(iterMul*math.Log2(float64(len(ins.VariableMap()))))
		for iteration := 0; iteration < numIterations; iteration++ {
			absoluteEtaChange, graph = iterateSurveyPropagationGraph(ins, graph, smooth)
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
			nonTrivialCover, variableOut, valueOut = surveyDecimation(ins, graph, smooth)
		}
	}
	return converged, nonTrivialCover, variableOut, valueOut
}
