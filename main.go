package main

import (
	"fmt"

	"github.com/khanhhhh/sat/guesser/maxmin"
	"github.com/khanhhhh/sat/guesser/surveydecimation"
	"github.com/khanhhhh/sat/instance"
	"github.com/khanhhhh/sat/solver/cdcl"
	"github.com/khanhhhh/sat/solver/search"
)

func main() {
	test2()
}

func test2() {
	ins := instance.Random3SAT(10, 3.2)
	{
		sat, assignment := cdcl.Solve(ins)
		eval, _ := ins.Evaluate(assignment)
		fmt.Println(sat, eval)
	}
	{
		guesser1 := func(ins instance.Instance) (variableOut instance.Variable, valueOut bool) {
			variableOut, valueOut = maxmin.Guess(ins, 1)
			return variableOut, valueOut
		}
		guesser2 := func(ins instance.Instance) (variableOut instance.Variable, valueOut bool) {
			var converged bool
			var nonTrivial bool
			converged, nonTrivial, variableOut, valueOut = surveydecimation.Guess(ins, 1.0)
			if converged == false || nonTrivial == false {
				variableOut, valueOut = maxmin.Guess(ins, 1)
			}
			return variableOut, valueOut
		}
		_ = guesser1
		_ = guesser2
		sat, assignment := search.Solve(ins, guesser1)
		eval, _ := ins.Evaluate(assignment)
		fmt.Println(sat, eval)
	}
}

func test1() {
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
				converged, nonTrivial, variable, value := surveydecimation.Guess(ins, 1.0)
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
					variable, value := maxmin.Guess(ins, 1.0)
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
