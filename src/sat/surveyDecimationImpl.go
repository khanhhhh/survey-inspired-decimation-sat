package sat

import "math"

var tolerance float64 = 0.001
var smooth float64 = 1.0

func (ins *instance) predict() (bool, map[variable]bool) {
	var etaChange float64 = 1
	g := ins.makePropagationGraph()
	numIterations := 1 + int(100*math.Log2(float64(len(ins.allVariables()))))
	iteration := 0
	for etaChange > tolerance && iteration < numIterations {
		iteration++
		etaChange, g = ins.iteratePropagationGraph(g, smooth)
	}
	if etaChange > tolerance {
		return false, nil
	}
	trivialCover, i, value := ins.decimation(g, smooth)
	if trivialCover {
		return ins.WalkSAT()
	}
	out := make(map[variable]bool)
	out[i] = value
	return true, out
}

func (ins *instance) SurveyInspiredDecimation() (bool, map[variable]bool) {
	numVariables := len(ins.allVariables())
	solution := make(map[variable]bool)
	for len(solution) < numVariables {
		sat, prediction := ins.predict()
		if sat == false {
			return false, nil
		}
		for i, value := range prediction {
			solution[i] = value
			ins.reduce(i, value)
			if ins.emptyClause() {
				panic("empty clause")
			}
		}
	}
	return true, solution
}
