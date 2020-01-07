package instance

import "math/rand"

type Variable = int
type Clause = int

// Instance :
// SAT Instance
type Instance interface {
	// basic
	Clone() (InstanceOut Instance)
	PushClause(variableMap map[Variable]bool)
	Evaluate(assignment map[Variable]bool) (sat bool, conflict Clause)
	// raw data
	VariableMap() (mapOut map[Variable]map[Clause]bool)
	ClauseMap() (mapOut map[Clause]map[Variable]bool)
}

// NewInstance :
// New empty SAT Instance
func NewInstance() (InstanceOut Instance) {
	return &instance{
		make(map[Variable]map[Clause]bool),
		make(map[Clause]map[Variable]bool),
	}
}

// Random3SAT :
// create a randomly generated 3-SAT Instance
func Random3SAT(numVariables int, density float64) (InstanceOut Instance) {
	numClauses := int(density * float64(numVariables))
	InstanceOut = NewInstance()
	for c := 0; c < numClauses; c++ {
		v1 := rand.Intn(numVariables)
		s1 := (rand.Intn(2) == 1)
		v2 := rand.Intn(numVariables)
		s2 := (rand.Intn(2) == 1)
		v3 := rand.Intn(numVariables)
		s3 := (rand.Intn(2) == 1)
		variableMap := make(map[Variable]bool)
		variableMap[v1] = s1
		variableMap[v2] = s2
		variableMap[v3] = s3
		InstanceOut.PushClause(variableMap)
	}
	return InstanceOut
}
