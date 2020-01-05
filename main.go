package main

import "sat"

import "math/rand"

import "fmt"

import "sync"

func main() {
	re := rand.New(rand.NewSource(1234154342))
	numTests := 1
	instance := make([]sat.Instance, numTests)
	for i := range instance {
		instance[i] = sat.Random3SAT(re, 40, 2.0)
	}
	var wg sync.WaitGroup
	for _, ins := range instance {
		wg.Add(1)
		go func(ins sat.Instance) {
			defer wg.Done()
			sat, assignment := ins.SurveyInspiredDecimation()
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
