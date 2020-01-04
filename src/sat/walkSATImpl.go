package sat

import "math/rand"

import "fmt"

var randFlip int = 100

func (ins *instance) WalkSAT() (bool, map[variable]bool) {
	out := make(map[variable]bool)
	for _, i := range ins.allVariables() {
		out[i] = (rand.Intn(2) == 1)
	}
	varList := make([]variable, 0)
	for i := range out {
		varList = append(varList, i)
	}
	numIterations := 100 * len(out)
	fmt.Printf("doing %v iterations walkSAT\n", numIterations)
	iteration := 0
	for iteration < numIterations {
		iteration++
		sat, conflictClause := ins.Evaluate(out)
		if sat {
			return true, out
		}
		if rand.Intn(randFlip) == 0 {
			i := varList[rand.Intn(len(varList))]
			out[i] = !out[i]
		} else {
			varList := make([]variable, 0)
			for i := range ins.clauseMap[conflictClause] {
				varList = append(varList, i)
			}
			i := varList[rand.Intn(len(varList))]
			out[i] = !out[i]
		}
	}
	return false, nil
}
