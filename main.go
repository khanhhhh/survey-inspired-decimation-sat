package main

import (
	"fmt"

	"github.com/khanhhhh/sat/guesser/surveydecimation"
	"github.com/khanhhhh/sat/instance"
	"github.com/khanhhhh/sat/solver/cdcl"
)

func main() {
	ins := instance.Random3SAT(40, 3.2)
	sat, assignment := cdcl.Solve(ins)
	eval, conflict := ins.Evaluate(assignment)
	fmt.Println(sat, eval, conflict)
	if sat {
		converged, nonTrivial, variable, value := surveydecimation.Guess(ins)
		if converged && nonTrivial {
			clause := make(map[instance.Variable]bool)
			clause[variable] = value
			ins.PushClause(clause)
			sat, _ := cdcl.Solve(ins)
			fmt.Println("prediction is: ", sat)
		}
	}
}
