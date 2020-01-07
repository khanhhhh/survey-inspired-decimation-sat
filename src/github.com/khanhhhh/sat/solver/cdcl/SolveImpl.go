package cdcl

import (
	"github.com/irifrance/gini"
	"github.com/irifrance/gini/z"
	"github.com/khanhhhh/sat/instance"
)

func makeVariableMap(mapIn map[instance.Variable]map[instance.Clause]bool) (mapOut map[instance.Variable]z.Var) {
	mapOut = make(map[instance.Variable]z.Var)
	counter := 0
	for variable := range mapIn {
		counter++
		mapOut[variable] = z.Var(counter)
	}
	return mapOut
}

// Solve :
// Solve SAT Instance exactly
func Solve(ins instance.Instance) (sat bool, assignment map[instance.Variable]bool) {
	variableMap := ins.VariableMap()
	clauseMap := ins.ClauseMap()
	var2var := makeVariableMap(variableMap)
	numVariables := len(variableMap)
	numClauses := len(clauseMap)
	g := gini.NewVc(numVariables, numClauses)
	// create gini SAT instance
	{
		for clause := range clauseMap {
			for variable, val := range clauseMap[clause] {
				if val {
					g.Add(var2var[variable].Pos())
				} else {
					g.Add(var2var[variable].Neg())
				}
			}
			g.Add(0)
		}
	}
	// solve
	{
		sat = (g.Solve() == 1)
		if sat {
			assignment = make(map[instance.Variable]bool)
			for variable := range variableMap {
				assignment[variable] = g.Value(var2var[variable].Pos())
			}
		}
	}
	return sat, assignment
}
