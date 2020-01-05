package main

import (
	"fmt"
	"sat"
)

func main() {
	converged := 0
	predicted := 0
	correct := 0
	incorrect := 0
	for iter := 1; ; iter++ {
		ins := sat.Random3SAT(40, 4.0)
		ok, nonTrivialCover, variable, value := ins.Predict()
		_ = ok
		_ = nonTrivialCover
		_ = variable
		_ = value
		if ok {
			converged++
		}
		if ok && nonTrivialCover {
			predicted++
			solved, assignment := ins.WalkSAT()
			if solved {
				if assignment[variable] == value {
					correct++
				} else {
					incorrect++
				}
			}
		}
		fmt.Printf("ITERATION %v\n", iter)
		convergedRate := float32(converged) / float32(iter)
		predictedRate := float32(predicted) / float32(iter)
		fmt.Printf("converged rate: %v\n", convergedRate)
		fmt.Printf("predicted rate: %v\n", predictedRate)
		upperPrecision := float32(predicted-incorrect) / float32(predicted)
		lowerPrecision := float32(correct) / float32(predicted)
		fmt.Printf("%.4f > precision > %.4f\n", upperPrecision, lowerPrecision)
		fmt.Printf("%.4f > combined precision > %.4f\n",
			upperPrecision*predictedRate+0.5*(1.0-predictedRate),
			lowerPrecision*predictedRate+0.5*(1.0-predictedRate),
		)
		fmt.Println()
	}
}
