package sat

type variable = int
type clause = int

type edge struct {
	variable variable
	clause   clause
}

func newEdge(variableIn variable, clauseIn clause) (edgeOut edge) {
	edgeOut = edge{variableIn, clauseIn}
	return edgeOut
}
