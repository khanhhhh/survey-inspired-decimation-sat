package sat

import "github.com/khanhhhh/sat/rational"

// surveyDecimation :
// inference max bias variable from a Survey Propagation Graph
func (ins *instance) surveyDecimation(graphIn *surveyPropagationGraph, smooth float64) (nonTrivialCover bool, maxBiasVariable variable, maxBiasValue bool) {
	var maxBias = rational.FromInt(0, 1)
	// select maxBias over all variables
	for _, variable := range ins.allVariables() {
		// calculate mu
		var mu [3]rational.Rat
		{
			var productPositive = rational.FromInt(1, 1)
			var productNegative = rational.FromInt(1, 1)
			for _, clause := range ins.clausePositive(variable) {
				productPositive = rational.Mul(
					productPositive,
					rational.Sub(rational.FromInt(1, 1), graphIn.etaMap[newEdge(variable, clause)]),
				)
			}
			for _, clause := range ins.clauseNegative(variable) {
				productNegative = rational.Mul(
					productNegative,
					rational.Sub(rational.FromInt(1, 1), graphIn.etaMap[newEdge(variable, clause)]),
				)
			}
			smoothConst := rational.FromFloat(smooth)
			mu[0] = rational.Mul(
				productPositive,
				rational.Sub(rational.FromInt(1, 1), rational.Mul(smoothConst, productNegative)),
			)

			mu[1] = rational.Mul(
				productNegative,
				rational.Sub(rational.FromInt(1, 1), rational.Mul(smoothConst, productPositive)),
			)

			mu[2] = rational.Mul(
				smoothConst,
				rational.Mul(productPositive, productNegative),
			)
		}
		// normalize
		{
			sum := rational.Add(rational.Add(mu[0], mu[1]), mu[2])
			if sum > 0 {
				mu[0] = rational.Div(mu[0], sum)
				mu[1] = rational.Div(mu[1], sum)
				mu[2] = rational.Div(mu[2], sum)
			}
		}
		// select maxBias
		{
			bias := rational.Abs(rational.Sub(mu[1], mu[0]))
			if bias > maxBias {
				maxBias = bias
				maxBiasVariable = variable
				maxBiasValue = (rational.Sub(mu[1], mu[0]).Sign() == 1)
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
