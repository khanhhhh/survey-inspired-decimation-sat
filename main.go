package main

import "sat"

import "math/rand"

import "fmt"

func main() {
	re := rand.New(rand.NewSource(1234154342))
	ins := sat.Random3SAT(re, 1024, 4.0)
	sat, assignment := ins.SurveyInspiredDecimation()
	if sat {
		eval, _ := ins.Evaluate(assignment)
		fmt.Println(sat, eval)
	} else {
		fmt.Println("failed")
	}
}
