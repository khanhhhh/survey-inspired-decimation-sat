package sat

import "math/rand"

func (ins *instance) walkSAT() map[variable]bool {
	out := make(map[variable]bool)
	for _, i := range ins.allVariables() {
		out[i] = (rand.Intn(2) == 1)
	}
	//iteration := 0
	return nil
}
