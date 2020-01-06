package sat

import "math/rand"

// Instance :
// SAT Instance
type Instance interface {
	Clone() (InstanceOut Instance)
	PushClause(...Literal)
	SidPredict() (converged bool, nonTrivialCover bool, variableOut variable, valueOut bool)
	MaxMinPredict() (variableOut variable, valueOut bool)
	WalkSolve() (sat bool, assignment map[variable]bool)
	CdclSolve() (sat bool, assignment map[variable]bool)
	Evaluate(assignment map[variable]bool) (sat bool, conflict clause)
}

// Literal :
// a pair {Index, Sign} corresponds to a literal
type Literal struct {
	Index int
	Sign  bool
}

// NewInstance :
// New empty SAT Instance
func NewInstance() (InstanceOut Instance) {
	return &instance{
		make(map[variable]map[clause]bool),
		make(map[clause]map[variable]bool),
	}
}

// Clone :
// Clone a SAT Instance
func (ins *instance) Clone() (InstanceOut Instance) {
	variableMap := make(map[variable]map[clause]bool)
	clauseMap := make(map[clause]map[variable]bool)
	// copy
	for i := range ins.variableMap {
		variableMap[i] = make(map[clause]bool)
		for a, value := range ins.variableMap[i] {
			variableMap[i][a] = value
		}
	}
	for a := range ins.clauseMap {
		clauseMap[a] = make(map[variable]bool)
		for i, value := range ins.clauseMap[a] {
			clauseMap[a][i] = value
		}
	}
	return &instance{variableMap, clauseMap}
}

// PushClause :
// Push a set of literals as a the clause
func (ins *instance) PushClause(someLiterals ...Literal) {
	nextClauseIndex := len(ins.clauseMap)
	// clauseMap
	{
		nextClause := make(map[variable]bool)
		for _, l := range someLiterals {
			variable := l.Index
			sign := l.Sign
			nextClause[variable] = sign
		}
		ins.clauseMap[nextClauseIndex] = nextClause
	}
	// variableMap
	{
		for _, l := range someLiterals {
			variable := l.Index
			sign := l.Sign
			_, exist := ins.variableMap[variable]
			if exist == false {
				ins.variableMap[variable] = make(map[clause]bool)
			}
			ins.variableMap[variable][nextClauseIndex] = sign
		}
	}
}

// Random3SAT :
// create a randomly generated 3-SAT Instance
func Random3SAT(numVariables int, density float64) (InstanceOut Instance) {
	numClauses := int(density * float64(numVariables))
	InstanceOut = NewInstance()
	for c := 0; c < numClauses; c++ {
		l1 := rand.Intn(numVariables)
		s1 := (rand.Intn(2) == 1)
		l2 := rand.Intn(numVariables)
		s2 := (rand.Intn(2) == 1)
		l3 := rand.Intn(numVariables)
		s3 := (rand.Intn(2) == 1)
		InstanceOut.PushClause(
			Literal{Index: l1, Sign: s1},
			Literal{Index: l2, Sign: s2},
			Literal{Index: l3, Sign: s3},
		)
	}
	return InstanceOut
}
