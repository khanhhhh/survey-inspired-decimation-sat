package walk

import (
	"math"
	"math/rand"

	"github.com/khanhhhh/sat/instance"
)

var randFlip int = 64
var base float64 = 1.4

// Solve :
// Solve SAT Instance stochastically
func Solve(ins instance.Instance) (sat bool, assignment map[instance.Variable]bool) {
	assignment = make(map[instance.Variable]bool)
	var variableList = make([]instance.Variable, 0)
	for variable := range ins.VariableMap() {
		variableList = append(variableList, variable)
	}
	var clauseMap = make(map[instance.Clause][]instance.Variable)
	for clause := range ins.ClauseMap() {
		variableSubList := make([]instance.Variable, 0)
		for variable := range ins.ClauseMap()[clause] {
			variableSubList = append(variableSubList, variable)
		}
		clauseMap[clause] = variableList
	}
	// randomize assignment
	for _, i := range variableList {
		assignment[i] = (rand.Intn(2) == 1)
	}
	// random instance.Variable
	randomVariable := func(variableList []instance.Variable) (variableOut instance.Variable) {
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
		// choose 1/randFlip probability to random flip any instance.Variable (overcome local optimum)
		// otherwise, flip the conflict clause instance.Variable
		var variableSubList []instance.Variable
		if rand.Intn(randFlip) == 0 {
			variableSubList = variableList
		} else {
			variableSubList = clauseMap[conflictClause]
		}
		chosenVariable := randomVariable(variableSubList)
		assignment[chosenVariable] = !assignment[chosenVariable]

	}
	// timeout
	return false, nil
}
