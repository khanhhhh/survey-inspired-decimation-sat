package sat

import "math"

import "fmt"

var tolerance float64 = 0.001
var smooth float64 = 1.0

func (ins *instance) predict() (variable, bool) {
	numIterations := 1 + int(math.Log2(float64(len(ins.allVariables()))))
	var etaChange float64 = 1
	g := ins.makePropagationGraph()
	iterations := 0
	for etaChange > tolerance && iterations < 1000*numIterations {
		iterations++
		etaChange, g = ins.iteratePropagationGraph(g, smooth)
	}
	if etaChange > tolerance {
		panic("halt")
	}
	trivialCover, variable, value := ins.decimation(g, smooth)
	if trivialCover {
		panic("trivial cover")
	}
	return variable, value
}

func (ins *instance) Solve() map[variable]bool {
	numVariables := len(ins.allVariables())
	solution := make(map[variable]bool)
	for len(solution) < numVariables {
		variable, value := ins.predict()
		solution[variable] = value
		fmt.Println(len(solution))
		ins.reduce(variable, value)
	}
	return solution
}
