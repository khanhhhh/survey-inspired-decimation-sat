package main

import (
	"fmt"
	"github.com/khanhhhh/sat"
)

func main() {
	satCount := 0
	convergedCountSID := 0
	nonTrivialCountSID := 0
	trueCountSID := 0
	trueCountMaxMin := 0
	for iter := 1; ; iter++ {
		ins := sat.Random3SAT(128, 4.2)
		isSat, _ := ins.CdclSolve()
		if isSat {
			satCount++
			{
				ok, nonTrivialCover, variable, value := ins.SidPredict()
				if ok {
					convergedCountSID++
				}
				if ok && nonTrivialCover {
					nonTrivialCountSID++
					// test
					ins.PushClause(sat.Literal{
						Index: variable,
						Sign:  value,
					})
					isSat, _ := ins.CdclSolve()
					if isSat {
						trueCountSID++
					}
				}
			}
			{
				variable, value := ins.MaxMinPredict()
				// test
				ins.PushClause(sat.Literal{
					Index: variable,
					Sign:  value,
				})
				isSat, _ := ins.CdclSolve()
				if isSat {
					trueCountMaxMin++
				}
			}
			fmt.Printf("Precision SID   :\t%.4f\n",
				float32(trueCountSID)/float32(nonTrivialCountSID),
			)
			fmt.Printf("Recall    SID   :\t%.4f\n",
				float32(trueCountSID)/float32(satCount),
			)
			fmt.Printf("Precision MaxMin:\t%.4f\n",
				float32(trueCountMaxMin)/float32(satCount),
			)
			fmt.Println()
		}
	}
}
