package sat

import (
	"fmt"
	"math"
)

var tolerance float64 = 0.001
var smooth float64 = 1.0

func (ins *instance) predict() (bool, map[variable]bool) {
	var etaChange float64 = 1
	g := ins.makePropagationGraph()
	numIterations := 1 + int(100*math.Log2(float64(len(ins.allVariables()))))
	iteration := 0
	for etaChange > tolerance && iteration < numIterations {
		iteration++
		etaChange, g = ins.iteratePropagationGraph(g, smooth)
	}
	if etaChange > tolerance {
		return false, nil
	}
	trivialCover, i, value := ins.decimation(g, smooth)
	if trivialCover {
		return ins.WalkSAT()
	}
	out := make(map[variable]bool)
	out[i] = value
	return true, out
}

func (ins *instance) SurveyInspiredDecimation() (bool, map[variable]bool) {
	sat, prediction := ins.predict()
	if sat == false {
		return false, nil
	}
	if len(prediction) == 1 {
		i := 0
		value := false
		for ii, ss := range prediction {
			i = ii
			value = ss
		}
		// ins1
		{
			ins1 := ins.clone()
			ins1.reduce(i, value)
			if ins1.emptyClause() {
				panic("empty clause")
			}
			sat, assignment := ins1.SurveyInspiredDecimation()
			if sat {
				return sat, assignment
			}
		}
		// ins2
		{
			ins2 := ins.clone()
			ins2.reduce(i, !value)
			if ins2.emptyClause() {
				//panic("empty clause")
				return false, nil
			}
			sat, assignment := ins2.SurveyInspiredDecimation()
			if sat {
				fmt.Println("prediction failed")
				return sat, assignment
			}
			return false, nil
		}
	} else {
		return true, prediction
	}
}
