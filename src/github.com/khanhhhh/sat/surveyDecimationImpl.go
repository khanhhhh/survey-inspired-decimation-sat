package sat

// surveyDecimation :
// inference max bias variable from a Survey Propagation Graph
func (ins *instance) surveyDecimation(graphIn *surveyPropagationGraph, smooth float64) (nonTrivialCover bool, maxBiasVariable variable, maxBiasValue bool) {
	var maxBias message = 0
	// select maxBias over all variables
	for _, variable := range ins.allVariables() {
		// calculate mu
		var mu [3]message
		{
			var productPositive message = 1
			var productNegative message = 1
			for _, clause := range ins.clausePositive(variable) {
				productPositive *= 1 - graphIn.etaMap[edge{variable, clause}]
			}
			for _, clause := range ins.clauseNegative(variable) {
				productNegative *= 1 - graphIn.etaMap[edge{variable, clause}]
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
			bias := absMessage(mu[1] - mu[0])
			if bias > maxBias {
				maxBias = bias
				maxBiasVariable = variable
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
