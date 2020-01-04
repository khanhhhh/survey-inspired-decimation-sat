package sat

type variable = int
type clause = int

type edge struct {
	variable variable
	clause   clause
}

type message = float64

func abs(m message) message {
	var out message
	if m >= 0 {
		out = +m
	} else {
		out = -m
	}
	return out
}
