package sat

func (ins *instance) surveyDecimation(graphIn *surveyPropagationGraph, smooth float64) (nonTrivialCover bool, maxBiasVariable variable, maxBiasValue bool) {
	var maxBias message = 0
	// select maxBias over all variables
	for _, i := range ins.allVariables() {
		// calculate mu
		var mu [3]message
		{
			var productPositive message = 1
			var productNegative message = 1
			for _, b := range ins.clausePositive(i) {
				productPositive *= 1 - graphIn.etaMap[edge{i, b}]
			}
			for _, b := range ins.clauseNegative(i) {
				productNegative *= 1 - graphIn.etaMap[edge{i, b}]
			}
			mu[0] = (1 - smooth*productNegative) * productPositive
			mu[1] = (1 - smooth*productPositive) * productNegative
			mu[2] = smooth * productNegative * productPositive
		}
		// normalize
		{
			sum := mu[0] + mu[1] + mu[2]
			if sum > 0 {
				mu[0] = mu[0] / sum
				mu[1] = mu[1] / sum
				mu[2] = mu[2] / sum
			}
		}
		// select maxBias
		{
			bias := abs(mu[1] - mu[0])
			if bias > maxBias {
				maxBias = bias
				maxBiasVariable = i
				maxBiasValue = (mu[1] > mu[0])
			}
		}
	}
	// detect trivial cover
	if maxBias == 0 {
		nonTrivialCover = false
	} else {
		nonTrivialCover = true
	}
	return nonTrivialCover, maxBiasVariable, maxBiasValue
}
