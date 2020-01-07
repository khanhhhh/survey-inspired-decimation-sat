package maxmin

import "github.com/khanhhhh/sat/instance"

func abs(valueIn float64) (valueOut float64) {
	if valueIn >= 0 {
		valueOut = +valueIn
	} else {
		valueOut = -valueIn
	}
	return valueOut
}

// Guess :
func Guess(ins instance.Instance, smooth float64) (variableOut instance.Variable, valueOut bool) {
	var maxBias float64 = -1
	variableMap := ins.VariableMap()
	for variable := range variableMap {
		var count float64 = 0
		for _, value := range variableMap[variable] {
			if value {
				count += 1.0
			} else {
				count -= smooth
			}
		}
		if abs(count) > maxBias {
			maxBias = abs(count)
			variableOut = variable
			valueOut = (count > 0)
		}
	}
	return variableOut, valueOut
}
