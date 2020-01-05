package sat

func (ins *instance) decimation(g *propagationGraph, smooth float64) (bool, variable, bool) {
	var nonTrivialCover bool
	var maxBias message = 0
	var maxBiasVariable variable
	var maxBiasSign bool
	for _, i := range ins.allVariables() {
		var productPositive message = 1
		var productNegative message = 1
		for _, b := range ins.clausePositive(i) {
			productPositive *= 1 - g.etaMap[edge{i, b}]
		}
		for _, b := range ins.clauseNegative(i) {
			productNegative *= 1 - g.etaMap[edge{i, b}]
		}
		var mu [3]message
		mu[0] = (1 - smooth*productNegative) * productPositive
		mu[1] = (1 - smooth*productPositive) * productNegative
		mu[2] = smooth * productNegative * productPositive
		sum := mu[0] + mu[1] + mu[2]
		if sum > 0 {
			mu[0] = mu[0] / sum
			mu[1] = mu[1] / sum
			mu[2] = mu[2] / sum
		}
		bias := abs(mu[1] - mu[0])
		if bias > maxBias {
			maxBias = bias
			maxBiasVariable = i
			maxBiasSign = (mu[1] > mu[0])
		}
	}
	if maxBias == 0 {
		nonTrivialCover = false
	} else {
		nonTrivialCover = true
	}
	return nonTrivialCover, maxBiasVariable, maxBiasSign
}
