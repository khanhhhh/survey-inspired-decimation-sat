package instance

type instance struct {
	variableMap map[Variable]map[Clause]bool
	clauseMap   map[Clause]map[Variable]bool
}

// Clone :
// Clone a SAT Instance
func (ins *instance) Clone() (InstanceOut Instance) {
	variableMap := make(map[Variable]map[Clause]bool)
	clauseMap := make(map[Clause]map[Variable]bool)
	// copy
	for i := range ins.variableMap {
		variableMap[i] = make(map[Clause]bool)
		for a, value := range ins.variableMap[i] {
			variableMap[i][a] = value
		}
	}
	for a := range ins.clauseMap {
		clauseMap[a] = make(map[Variable]bool)
		for i, value := range ins.clauseMap[a] {
			clauseMap[a][i] = value
		}
	}
	return &instance{variableMap, clauseMap}
}

// PushClause :
// Push a set of literals as a the Clause
func (ins *instance) PushClause(variableMap map[Variable]bool) {
	nextClauseIndex := len(ins.clauseMap)
	// clauseMap
	{
		nextClause := make(map[Variable]bool)
		for variable, value := range variableMap {
			nextClause[variable] = value
		}
		ins.clauseMap[nextClauseIndex] = nextClause
	}
	// variableMap
	{
		for variable, value := range variableMap {
			_, exist := ins.variableMap[variable]
			if exist == false {
				ins.variableMap[variable] = make(map[Clause]bool)
			}
			ins.variableMap[variable][nextClauseIndex] = value
		}
	}
}

func (ins *instance) VariableMap() (mapOut map[Variable]map[Clause]bool) {
	mapOut = ins.variableMap
	return mapOut
}
func (ins *instance) ClauseMap() (mapOut map[Clause]map[Variable]bool) {
	mapOut = ins.clauseMap
	return mapOut
}

/*
// hasEmptyClause :
// return true if there are some empty Clause
func (ins *instance) hasEmptyClause() bool {
	for c := range ins.clauseMap {
		if len(ins.clauseMap[c]) == 0 {
			return true
		}
	}
	return false
}
// reduce :
// reduce the assignment by setting value to a Variable
func (ins *instance) reduce(i Variable, value bool) {
	// variableMap
	delete(ins.variableMap, i)
	// clauseMap
	for c := range ins.clauseMap {
		sign, exist := ins.clauseMap[c][i]
		if exist {
			if sign == value { // remove the Clause
				delete(ins.clauseMap, c)
			} else { // reduce the Clause
				delete(ins.clauseMap[c], i)
			}
		}
	}
}
*/
