package main

import (
	"fmt"
	"time"

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
	var durationSID time.Duration = 0
	var durationSolver time.Duration = 0
	solver := cdcl.Solve
	for iter := 1; ; iter++ {
		ins := instance.Random3SAT(128, 4.1)
		t := time.Now()
		sat, assignment := solver(ins)
		dt := time.Since(t)
		eval, _ := ins.Evaluate(assignment)
		if sat {
			durationSolver += dt
			if eval == false {
				panic("cdcl failed")
			}
			counter++
			{
				ins := ins.Clone()
				t := time.Now()
				converged, nonTrivial, variable, value := surveydecimation.Guess(ins, 1.0)
				dt := time.Since(t)
				durationSID += dt
				if converged && nonTrivial {
					convergentCounterSID++
					// test
					ins.Reduce(variable, value)
					sat, _ := solver(ins)
					if sat {
						trueCounterSID++
					}
				}

				convergentRate := float32(convergentCounterSID) / float32(counter)
				sidPrecision := float32(trueCounterSID) / float32(convergentCounterSID)
				effSidPrecision := convergentRate*sidPrecision + (1-convergentRate)*0.5
				durationSidAvg := time.Duration(int(durationSID) / counter)
				durationSolverAvg := time.Duration(int(durationSolver) / counter)
				fmt.Printf("Iter: %v/%v\n", counter, iter)
				fmt.Printf("Sover duration: %v\n", durationSolverAvg)
				fmt.Printf("SID convergent rate: %.4f\n", convergentRate)
				fmt.Printf("SID precision:       %.4f\n", sidPrecision)
				fmt.Printf("SID eff precision:   %.4f\n", effSidPrecision)
				fmt.Printf("SID duration:        %v\n", durationSidAvg)
				fmt.Println()
			}
		}
	}
}
