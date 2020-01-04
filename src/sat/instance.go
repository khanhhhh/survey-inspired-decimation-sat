package sat

import "math/rand"

// Instance :
type Instance interface {
	PushClause(...Literal)
	Solve() map[variable]bool
}

// Literal :
type Literal struct {
	Index int
	Sign  bool
}

// NewInstance :
func NewInstance() Instance {
	return &instance{
		make(map[variable]map[clause]bool),
		make(map[clause]map[variable]bool),
	}
}

// PushClause :
func (ins *instance) PushClause(someLiterals ...Literal) {
	nextClauseIndex := len(ins.clauseMap)
	// clauseMap
	nextClause := make(map[variable]bool)
	for _, l := range someLiterals {
		variable := l.Index
		sign := l.Sign
		nextClause[variable] = sign
	}
	ins.clauseMap[nextClauseIndex] = nextClause
	// variableMap
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

// Random3SAT :
func Random3SAT(rand *rand.Rand, numVariables int, density float64) Instance {
	numClauses := int(density * float64(numVariables))
	out := NewInstance()
	for c := 0; c < numClauses; c++ {
		l1 := rand.Intn(numVariables)
		s1 := (rand.Intn(2) == 1)
		l2 := rand.Intn(numVariables)
		s2 := (rand.Intn(2) == 1)
		l3 := rand.Intn(numVariables)
		s3 := (rand.Intn(2) == 1)
		out.PushClause(
			Literal{Index: l1, Sign: s1},
			Literal{Index: l2, Sign: s2},
			Literal{Index: l3, Sign: s3},
		)
	}
	return out
}
