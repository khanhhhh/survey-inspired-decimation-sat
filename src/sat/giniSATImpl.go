package sat

import (
	"github.com/irifrance/gini"
	"github.com/irifrance/gini/z"
)

func (ins *instance) CdclSAT() (bool, map[variable]bool) {
	numVariables := ins.capVariables()
	numClauses := ins.capClauses()
	g := gini.NewVc(numVariables, numClauses)
	for a := range ins.clauseMap {
		for i, val := range ins.clauseMap[a] {
			if val {
				g.Add(z.Var(i + 1).Pos())
			} else {
				g.Add(z.Var(i + 1).Neg())
			}
		}
		g.Add(0)
	}
	var sat bool = (g.Solve() == 1)
	var assignment map[variable]bool = nil
	if sat {
		assignment = make(map[variable]bool)
		for i := range ins.variableMap {
			assignment[i] = g.Value(z.Var(i + 1).Pos())
		}
	}
	return sat, assignment
}
