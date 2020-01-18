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
	/*
		var cpuprofile = flag.String("cpuprofile", "./cpu.prof", "write cpu profile to `file`")
		var memprofile = flag.String("memprofile", "./mem.prof", "write memory profile to `file`")

		flag.Parse()
		if *cpuprofile != "" {
			f, err := os.Create(*cpuprofile)
			if err != nil {
				log.Fatal("could not create CPU profile: ", err)
			}
			defer f.Close()
			if err := pprof.StartCPUProfile(f); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
			defer pprof.StopCPUProfile()
		}

		// ... rest of the program ...
		defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	*/
	test()
	/*
		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal("could not create memory profile: ", err)
			}
			defer f.Close()
			runtime.GC() // get up-to-date statistics
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("could not write memory profile: ", err)
			}
		}
	*/
}

func test() {
	test4()
}

func test4() {
	for {
		ins := instance.Random3SATHard(40, cdcl.Solve)
		sat, _ := cdcl.Solve(ins)
		fmt.Println(sat)
	}
}

func test3() {
	ins := instance.Random3SAT(300, 4.0)
	completeSearch := cdcl.Solve
	t := time.Now()
	sat, assignment := completeSearch(ins)
	dt := time.Since(t)
	eval, _ := ins.Evaluate(assignment)
	fmt.Println("cdcl search:", sat, eval, dt)
	if sat && eval {
		for len(ins.VariableMap()) > 0 {
			fmt.Printf("predicting variable: %v\t", len(ins.VariableMap()))
			t := time.Now()
			converged, nonTrivial, variableOut, valueOut := surveydecimation.Guess(ins, 1)
			dt := time.Since(t)
			if converged && nonTrivial {
				// test
				ins.Reduce(variableOut, valueOut)
				if sat, _ := completeSearch(ins); sat == false {
					fmt.Println("failed: unsat", dt)
					break
				} else {
					fmt.Println("succeeded", dt)
				}
			} else {
				fmt.Println("failed: unconverged", dt)
				break
			}
		}
	}
}

func test2() {
	for {
		ins := instance.Random3SAT(500, 4.0)
		{
			t := time.Now()
			sat, assignment := surveysearch.Solve(ins)
			eval, _ := ins.Evaluate(assignment)
			fmt.Println("survey search:", sat, eval, time.Since(t))
		}
		{
			t := time.Now()
			sat, assignment := cdcl.Solve(ins)
			eval, _ := ins.Evaluate(assignment)
			fmt.Println("cdcl   search:", sat, eval, time.Since(t))
		}
		fmt.Println()
	}
}

func test1() {
	counter := 0
	convergentCounterSID := 0
	trueCounterSID := 0
	var durationSID time.Duration = 0
	var durationSolver time.Duration = 0
	solver := cdcl.Solve
	for iter := 1; iter < 100; iter++ {
		ins := instance.Random3SAT(200, 4.1)
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
