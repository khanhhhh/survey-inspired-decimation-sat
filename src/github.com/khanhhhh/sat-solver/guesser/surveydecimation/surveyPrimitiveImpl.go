package surveydecimation

import "github.com/khanhhhh/sat-solver/instance"

type edge struct {
	variable instance.Variable
	clause   instance.Clause
}

func newEdge(variableIn instance.Variable, clauseIn instance.Clause) (edgeOut edge) {
	edgeOut = edge{variableIn, clauseIn}
	return edgeOut
}

// C+(i) :
// set of clauses contain i as a positive literal
func clausePositive(ins instance.Instance, variable instance.Variable) (setOut []instance.Clause) {
	setOut = make([]instance.Clause, 0)
	for clause, sign := range ins.VariableMap()[variable] {
		if sign == true {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}

// C-(i) :
// set of clauses contain i as a negative literal
func clauseNegative(ins instance.Instance, variable instance.Variable) (setOut []instance.Clause) {
	setOut = make([]instance.Clause, 0)
	for clause, sign := range ins.VariableMap()[variable] {
		if sign == false {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}

// Cs(a, i) :
// set of clauses contain the same sign literal
func clauseAgree(ins instance.Instance, edgeIn edge) (setOut []instance.Clause) {
	setOut = make([]instance.Clause, 0)
	edgeClause := edgeIn.clause
	edgeVariable := edgeIn.variable
	edgeSign := ins.VariableMap()[edgeVariable][edgeClause]
	for clause, sign := range ins.VariableMap()[edgeVariable] {
		if clause != edgeClause && sign == edgeSign {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}

// Cu(a, i) :
// set of clauses contain the different sign literal
func clauseDisagree(ins instance.Instance, edgeIn edge) (setOut []instance.Clause) {
	setOut = make([]instance.Clause, 0)
	edgeClause := edgeIn.clause
	edgeVariable := edgeIn.variable
	edgeSign := ins.VariableMap()[edgeVariable][edgeClause]
	for clause, sign := range ins.VariableMap()[edgeVariable] {
		if clause != edgeClause && sign != edgeSign {
			setOut = append(setOut, clause)
		}
	}
	return setOut
}

// allEdges :
// list all edges of a instance
func allEdges(ins instance.Instance) (edgeOut []edge) {
	edgeOut = make([]edge, 0)
	for variable, clauseMap := range ins.VariableMap() {
		for clause := range clauseMap {
			edgeOut = append(edgeOut, newEdge(variable, clause))
		}
	}
	return edgeOut
}
