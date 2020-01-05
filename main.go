package main

import "sat"

import "math/rand"

import "fmt"

import "sync"

func main() {
	rand.Seed(1234154342)
	numTests := 1
	instance := make([]sat.Instance, numTests)
	for i := range instance {
		instance[i] = sat.Random3SAT(40, 4.2)
	}
	var wg sync.WaitGroup
	for _, ins := range instance {
		wg.Add(1)
		go func(ins sat.Instance) {
			defer wg.Done()
			//sat, assignment := ins.SurveyInspiredDecimation()
			sat, assignment := ins.WalkSAT()
			if sat {
				eval, _ := ins.Evaluate(assignment)
				fmt.Println(sat, eval)
			} else {
				fmt.Println("failed")
			}
		}(ins)
	}
	wg.Wait()
}
