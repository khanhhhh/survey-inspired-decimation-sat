package sat

import (
	"math"
)

var tolerance float64 = 0.001
var smooth float64 = 1.0

func (ins *instance) Predict() (bool, bool, variable, bool) {
	var converged bool
	var nonTrivialCover bool
	var variable variable
	var value bool
	{
		// survey propagation
		var etaChange float64 = 1
		g := ins.makePropagationGraph()
		{
			numIterations := 1 + int(100*math.Log2(float64(len(ins.allVariables()))))
			iteration := 0
			for etaChange > tolerance && iteration < numIterations {
				iteration++
				etaChange, g = ins.iteratePropagationGraph(g, smooth)
			}
		}
		// converge ?
		if etaChange > tolerance {
			converged = false
		} else {
			converged = true
			nonTrivialCover, variable, value = ins.decimation(g, smooth)
			if !nonTrivialCover {
				panic("trivial cover")
			}
		}
	}
	return converged, nonTrivialCover, variable, value
}
