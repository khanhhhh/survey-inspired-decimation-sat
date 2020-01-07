package main

import (
	"fmt"
	"time"

	"github.com/khanhhhh/sat/guesser/maxmin"
	"github.com/khanhhhh/sat/guesser/surveydecimation"
	"github.com/khanhhhh/sat/instance"
	"github.com/khanhhhh/sat/solver/cdcl"
	"github.com/khanhhhh/sat/solver/surveysearch"
)

func main() {
	test1()
}

func test2() {
	ins := instance.Random3SAT(512, 4.0)
	{
		t := time.Now()
		sat, assignment := surveysearch.Solve(ins)
		eval, _ := ins.Evaluate(assignment)
		fmt.Println(sat, eval, time.Since(t))
	}
	{
		t := time.Now()
		sat, assignment := cdcl.Solve(ins)
		eval, _ := ins.Evaluate(assignment)
		fmt.Println(sat, eval, time.Since(t))
	}
}

func test1() {
	counter := 0
	convergentCounterSID := 0
	trueCounterSID := 0
	trueCounterMaxMin := 0
	for {
		ins := instance.Random3SAT(128, 4.26)
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

				convergentRate := float32(convergentCounterSID) / float32(counter)
				sidPrecision := float32(trueCounterSID) / float32(convergentCounterSID)
				effSidPrecision := convergentRate*sidPrecision + (1-convergentRate)*0.5
				maxminPrecision := float32(trueCounterMaxMin) / float32(counter)

				fmt.Printf("SID convergent rate: %.4f\n", convergentRate)
				fmt.Printf("SID precision:       %.4f\n", sidPrecision)
				fmt.Printf("SID eff precision:   %.4f\n", effSidPrecision)
				fmt.Printf("MaxMin precision:    %.4f\n", maxminPrecision)
				fmt.Println()
			}
		}
	}
}
