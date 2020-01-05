package sat

type instance struct {
	variableMap map[variable]map[clause]bool
	clauseMap   map[clause]map[variable]bool
}

// variables
func (ins *instance) allVariables() []variable {
	out := make([]variable, 0)
	for i := range ins.variableMap {
		out = append(out, i)
	}
	return out
}

// edges
func (ins *instance) allEdges() []edge {
	out := make([]edge, 0)
	for v := range ins.variableMap {
		for c := range ins.variableMap[v] {
			out = append(out,
				edge{v, c},
			)
		}
	}
	return out
}

// clone
func (ins *instance) clone() *instance {
	variableMap := make(map[variable]map[clause]bool)
	clauseMap := make(map[clause]map[variable]bool)
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

// emptyClause :
// return true if there are some empty clause
func (ins *instance) emptyClause() bool {
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

// C+(i)
func (ins *instance) clausePositive(i variable) []clause {
	out := make([]clause, 0)
	for c, sign := range ins.variableMap[i] {
		if sign == true {
			out = append(out, c)
		}
	}
	return out
}

// C-(i)
func (ins *instance) clauseNegative(i variable) []clause {
	out := make([]clause, 0)
	for c, sign := range ins.variableMap[i] {
		if sign == false {
			out = append(out, c)
		}
	}
	return out
}

// Cs(a, i)
func (ins *instance) clauseAgree(e edge) []clause {
	a := e.clause
	i := e.variable
	out := make([]clause, 0)
	Jai := ins.variableMap[i][a]
	for c, sign := range ins.variableMap[i] {
		if c != a && sign == Jai {
			out = append(out, c)
		}
	}
	return out
}

// Cu(a, i)
func (ins *instance) clauseDisagree(e edge) []clause {
	a := e.clause
	i := e.variable
	out := make([]clause, 0)
	Jai := ins.variableMap[i][a]
	for c, sign := range ins.variableMap[i] {
		if c != a && sign != Jai {
			out = append(out, c)
		}
	}
	return out
}
