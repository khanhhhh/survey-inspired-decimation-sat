package sat

// IsComplete :
// return true if the assignment is complete
func (ins *instance) IsComplete(assignment map[variable]bool) (isComplete bool) {
	isComplete = true
	for variable := range ins.variableMap {
		_, exist := assignment[variable]
		if exist == false {
			isComplete = false
			break
		}
	}
	return isComplete
}

// Evaluate :
// Evaluate an partial assignment
// {
// return true || return false and the first conflict clause
// }
func (ins *instance) Evaluate(assignment map[variable]bool) (sat bool, conflict clause) {
	sat = true
	for conflictClause := range ins.clauseMap {
		unsat := 0
		for variable := range ins.clauseMap[conflictClause] {
			val, exist := assignment[variable]
			if exist && ins.clauseMap[conflictClause][variable] != val {
				unsat++
			}
		}
		if unsat == len(ins.clauseMap[conflictClause]) {
			sat = false
			conflict = conflictClause
			break
		}
	}
	return sat, conflict
}
