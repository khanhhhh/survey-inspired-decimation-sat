package sat

import "math/rand"

import "math"

var randFlip int = 64
var base float64 = 1.4

// WalkSolve :
// Solve SAT Instance stochastically
func (ins *instance) WalkSolve() (sat bool, assignment map[variable]bool) {
	assignment = make(map[variable]bool)
	allVariables := ins.allVariables()
	// randomize assignment
	for _, i := range allVariables {
		assignment[i] = (rand.Intn(2) == 1)
	}
	// random variable
	toList := func(variableSet map[variable]bool) (variableList []variable) {
		variableList = make([]variable, 0)
		for variable := range variableSet {
			variableList = append(variableList, variable)
		}
		return variableList
	}
	randomVariable := func(variableList []variable) (variableOut variable) {
		length := len(variableList)
		index := rand.Intn(length)
		variableOut = variableList[index]
		return variableOut
	}
	// loop
	numIterations := 1 + int(math.Pow(base, float64(len(assignment))))
	for iteration := 0; iteration < numIterations; iteration++ {
		sat, conflictClause := ins.Evaluate(assignment)
		// return if sat
		if sat {
			return true, assignment
		}
		// choose 1/randFlip probability to random flip any variable (overcome local optimum)
		// otherwise, flip the conflict clause variable
		if rand.Intn(randFlip) == 0 {
			variable := randomVariable(allVariables)
			assignment[variable] = !assignment[variable]
		} else {
			variableSet := ins.clauseMap[conflictClause]
			variableList := toList(variableSet)
			variable := randomVariable(variableList)
			assignment[variable] = !assignment[variable]
		}
	}
	// timeout
	return false, nil
}
