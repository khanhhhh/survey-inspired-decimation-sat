package sat

import "fmt"

var tolerance float64 = 0.01
var smooth float64 = 1.0

func (ins *instance) predict() (variable, bool) {
	var converged bool = false
	g := ins.makePropagationGraph()
	iterations := 0
	for converged == false {
		iterations++
		converged, g = ins.iteratePropagationGraph(g, smooth, tolerance)
	}
	trivialCover, variable, value := ins.decimation(g, smooth)
	if trivialCover {
		panic("trivial cover")
	}
	fmt.Printf("variable: %v, %v | iterations: %v\n", variable, value, iterations)
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
