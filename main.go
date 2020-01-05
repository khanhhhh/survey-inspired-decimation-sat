package main

import (
	"fmt"
	"sat"
)

func main() {
	correct := 0
	incorrect := 0
	predicted := 0
	for iter := 0; ; iter++ {
		ins := sat.Random3SAT(64, 3.0)
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
		//fmt.Printf("{predicted: %v, correct: %v, incorrect: %v}\t", predicted, correct, incorrect)
		fmt.Printf(
			"%v: %.2f > true probability > %.2f\n",
			iter,
			1-float32(incorrect)/float32(predicted),
			float32(correct)/float32(predicted),
		)
	}
}
