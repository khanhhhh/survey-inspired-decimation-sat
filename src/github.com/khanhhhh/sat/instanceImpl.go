package sat

type instance struct {
	variableMap map[variable]map[clause]bool
	clauseMap   map[clause]map[variable]bool
}

// allVariables :
// set of all variables
func (ins *instance) allVariables() (setOut []variable) {
	setOut = make([]variable, 0)
	for variable := range ins.variableMap {
		setOut = append(setOut, variable)
	}
	return setOut
}

// allClauses :
// set of all clauses
func (ins *instance) allClauses() (setOut []clause) {
	setOut = make([]clause, 0)
	for clause := range ins.clauseMap {
		setOut = append(setOut, clause)
	}
	return setOut
}

// allEdges :
// set of all edges
func (ins *instance) allEdges() (setOut []edge) {
	setOut = make([]edge, 0)
	for variable := range ins.variableMap {
		for clause := range ins.variableMap[variable] {
			setOut = append(setOut, edge{variable, clause})
		}
	}
	return setOut
}

// capVariables :
// range of all variable indices
func (ins *instance) capVariables() (maxVarIndex int) {
	maxVarIndex = 0
	for _, variable := range ins.allVariables() {
		if variable > maxVarIndex {
			maxVarIndex = variable
		}
	}
	return 1 + maxVarIndex
}

// capClauses :
// range of all clause indices
func (ins *instance) capClauses() (maxClauseIndex int) {
	maxClauseIndex = 0
	for _, clause := range ins.allClauses() {
		if clause > maxClauseIndex {
			maxClauseIndex = clause
		}
	}
	return maxClauseIndex
}

/*
// hasEmptyClause :
// return true if there are some empty clause
func (ins *instance) hasEmptyClause() bool {
	for c := range ins.clauseMap {
		if len(ins.clauseMap[c]) == 0 {
			return true
		}
	}
	return false
}
// reduce :
// reduce the assignment by setting value to a variable
func (ins *instance) reduce(i variable, value bool) {
	// variableMap
	delete(ins.variableMap, i)
	// clauseMap
	for c := range ins.clauseMap {
		sign, exist := ins.clauseMap[c][i]
		if exist {
			if sign == value { // remove the clause
				delete(ins.clauseMap, c)
			} else { // reduce the clause
				delete(ins.clauseMap[c], i)
			}
		}
	}
}
*/
