package main

import (
	"fmt"

	"github.com/khanhhhh/sat"
)

func main() {
	satCount := 0
	convergedCount := 0
	nonTrivialCount := 0
	trueCount := 0
	for iter := 1; ; iter++ {
		ins := sat.Random3SAT(128, 4.0)
		isSat, _ := ins.CdclSolve()
		if isSat {
			satCount++
			ok, nonTrivialCover, variable, value := ins.SidPredict()
			if ok {
				convergedCount++
			}
			if ok && nonTrivialCover {
				nonTrivialCount++
				// test
				ins.PushClause(sat.Literal{
					Index: variable,
					Sign:  value,
				})
				isSat, _ := ins.CdclSolve()
				if isSat {
					trueCount++
				}
			}
		}
		fmt.Println(iter, satCount, convergedCount, nonTrivialCount, trueCount)
	}
}
