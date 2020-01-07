package main

import (
	"fmt"
	"github.com/khanhhhh/sat/guesser/maxmin"
	"github.com/khanhhhh/sat/guesser/surveydecimation"
	"github.com/khanhhhh/sat/instance"
	"github.com/khanhhhh/sat/solver/cdcl"
)

func main() {
	counter := 0
	convergentCounterSID := 0
	trueCounterSID := 0
	trueCounterMaxMin := 0
	for {
		ins := instance.Random3SAT(128, 4.2)
		sat, assignment := cdcl.Solve(ins)
		eval, _ := ins.Evaluate(assignment)
		if sat {
			if eval == false {
				panic("cdcl failed")
			}
			counter++
			{
				ins := ins.Clone()
				converged, nonTrivial, variable, value := surveydecimation.Guess(ins)
				if converged && nonTrivial {
					convergentCounterSID++
					// test
					clause := make(map[instance.Variable]bool)
					clause[variable] = value
					ins.PushClause(clause)
					sat, _ := cdcl.Solve(ins)
					if sat {
						trueCounterSID++
					}
				}
				{
					ins := ins.Clone()
					variable, value := maxmin.Guess(ins)
					// test
					clause := make(map[instance.Variable]bool)
					clause[variable] = value
					ins.PushClause(clause)
					sat, _ := cdcl.Solve(ins)
					if sat {
						trueCounterMaxMin++
					}
				}

				fmt.Printf("SID convergent rate: %.4f\n", float32(convergentCounterSID)/float32(counter))
				fmt.Printf("SID precision:       %.4f\n", float32(trueCounterSID)/float32(convergentCounterSID))
				fmt.Printf("MaxMin precision:    %.4f\n", float32(trueCounterMaxMin)/float32(counter))
				fmt.Println()
			}
		}
	}
}
