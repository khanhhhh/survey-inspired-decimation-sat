package sat

type variable = int
type clause = int

type edge struct {
	variable variable
	clause   clause
}
