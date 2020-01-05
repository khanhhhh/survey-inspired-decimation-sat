package main

import (
	"fmt"
	"math/rand"
	"sat"
)

func main() {
	rand.Seed(1234154342)
	correct := 0
	incorrect := 0
	predicted := 0
	for {
		ins := sat.Random3SAT(256, 3.0)
		ok, variable, value := ins.Predict()
		if ok {
			predicted++
		}
		solved, assignment := ins.WalkSAT()
		if solved {
			if ok {
				if assignment[variable] == value {
					correct++
				} else {
					incorrect++
				}
			}
		}
		fmt.Printf("{predicted: %v, correct: %v, incorrect: %v}\n", predicted, correct, incorrect)
		fmt.Printf(
			"\t%v > true probability > %v\n",
			1-float32(incorrect)/float32(predicted),
			float32(correct)/float32(predicted),
		)
	}
}
