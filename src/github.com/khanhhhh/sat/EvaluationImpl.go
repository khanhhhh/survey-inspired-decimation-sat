package sat

func (ins *instance) Evaluate(assignment map[variable]bool) (bool, clause) {
	for a := range ins.clauseMap {
		unsat := 0
		for i := range ins.clauseMap[a] {
			val, exist := assignment[i]
			if exist && ins.clauseMap[a][i] != val {
				unsat++
			}
		}
		if unsat == len(ins.clauseMap[a]) {
			return false, a
		}
	}
	return true, 0
}
