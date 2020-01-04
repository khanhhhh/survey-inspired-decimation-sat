package sat

import "fmt"

import "math"

var tolerance float64 = 0.01
var smooth float64 = 1.0

func (ins *instance) predict() (variable, bool) {
	numIterations := 1 + int(math.Log2(float64(len(ins.allVariables()))))
	var etaChange float64 = 1
	g := ins.makePropagationGraph()
	iterations := 0
	for etaChange > tolerance && iterations < 100*numIterations {
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
	fmt.Printf("variable: %v, %v \t iterations: %v\n", variable, value, iterations)
	return variable, value
}

func (ins *instance) Solve() []bool {
	numVariables := len(ins.allVariables())
	solution := make(map[variable]bool)
	for len(solution) < numVariables {
		variable, value := ins.predict()
		solution[variable] = value
		ins.reduce(variable, value)
	}
	out := make([]bool, numVariables)
	for i := range out {
		out[i] = solution[i]
	}
	return out
}
