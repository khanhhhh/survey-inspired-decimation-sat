package maxmin

import "github.com/khanhhhh/sat-solver/instance"

func abs(valueIn int) (valueOut int) {
	if valueIn >= 0 {
		valueOut = +valueIn
	} else {
		valueOut = -valueIn
	}
	return valueOut
}

func Guess(ins instance.Instance) (variableOut instance.Variable, valueOut bool) {
	maxBias := -1
	variableMap := ins.VariableMap()
	for variable := range variableMap {
		count := 0
		for _, value := range variableMap[variable] {
			if value {
				count++
			} else {
				count--
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
