package sat


// C+(i) :
// set of clauses contain i as a positive literal
func (ins *instance) clausePositive(i variable) (setOut []clause) {
	setOut = make([]clause, 0)
	for clause, sign := range ins.variableMap[i] {
		if sign == true {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}

// C-(i) :
// set of clauses contain i as a negative literal
func (ins *instance) clauseNegative(i variable) (setOut []clause) {
	setOut = make([]clause, 0)
	for clause, sign := range ins.variableMap[i] {
		if sign == false {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}

// Cs(a, i) :
// set of clauses contain the same sign literal
func (ins *instance) clauseAgree(edgeIn edge) (setOut []clause) {
	setOut = make([]clause, 0)
	edgeClause := edgeIn.clause
	edgeVariable := edgeIn.variable
	edgeSign := ins.variableMap[edgeVariable][edgeClause]
	for clause, sign := range ins.variableMap[edgeVariable] {
		if clause != edgeClause && sign == edgeSign {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}

// Cu(a, i) :
// set of clauses contain the different sign literal
func (ins *instance) clauseDisagree(edgeIn edge) (setOut []clause) {
	setOut = make([]clause, 0)
	edgeClause := edgeIn.clause
	edgeVariable := edgeIn.variable
	edgeSign := ins.variableMap[edgeVariable][edgeClause]
	for clause, sign := range ins.variableMap[edgeVariable] {
		if clause != edgeClause && sign != edgeSign {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}
