package surveydecimation

import (
	"github.com/khanhhhh/sat/guesser/surveydecimation/message"
	"github.com/khanhhhh/sat/instance"
)

// surveyDecimation :
// inference max bias variable from a Survey Propagation Graph
func surveyDecimation(ins instance.Instance, graphIn *surveyPropagationGraph, smooth float64) (nonTrivialCover bool, maxBiasVariable instance.Variable, maxBiasValue bool) {
	var maxBias = message.FromInt(0, 1)
	// select maxBias over all variables
	for variable := range ins.VariableMap() {
		// calculate mu
		var mu [3]message.Message
		{
			var productPositive = message.FromInt(1, 1)
			var productNegative = message.FromInt(1, 1)
			for _, clause := range clausePositive(ins, variable) {
				productPositive = message.Mul(
					productPositive,
					message.Sub(message.FromInt(1, 1), graphIn.etaMap[newEdge(variable, clause)]),
				)
			}
			for _, clause := range clauseNegative(ins, variable) {
				productNegative = message.Mul(
					productNegative,
					message.Sub(message.FromInt(1, 1), graphIn.etaMap[newEdge(variable, clause)]),
				)
			}
			smoothConst := message.FromFloat(smooth)
			mu[0] = message.Mul(
				productPositive,
				message.Sub(message.FromInt(1, 1), message.Mul(smoothConst, productNegative)),
			)

			mu[1] = message.Mul(
				productNegative,
				message.Sub(message.FromInt(1, 1), message.Mul(smoothConst, productPositive)),
			)

			mu[2] = message.Mul(
				smoothConst,
				message.Mul(productPositive, productNegative),
			)
		}
		// normalize
		{
			sum := message.Add(message.Add(mu[0], mu[1]), mu[2])
			if sum > 0 {
				mu[0] = message.Div(mu[0], sum)
				mu[1] = message.Div(mu[1], sum)
				mu[2] = message.Div(mu[2], sum)
			}
		}
		// select maxBias
		{
			bias := message.Abs(message.Sub(mu[1], mu[0]))
			if bias > maxBias {
				maxBias = bias
				maxBiasVariable = variable
				maxBiasValue = (message.Sub(mu[1], mu[0]).Sign() == 1)
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
