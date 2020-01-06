package sat

func absInt(valueIn int) (valueOut int) {
	if valueIn >= 0 {
		valueOut = +valueIn
	} else {
		valueOut = -valueIn
	}
	return valueOut
}

func (ins *instance) MaxMinPredict() (variableOut variable, valueOut bool) {
	maxBias := -1
	for variable := range ins.variableMap {
		count := 0
		for _, value := range ins.variableMap[variable] {
			if value {
				count++
			} else {
				count--
			}
		}
		if absInt(count) > maxBias {
			maxBias = absInt(count)
			variableOut = variable
			valueOut = (count > 0)
		}
	}
	return variableOut, valueOut
}
