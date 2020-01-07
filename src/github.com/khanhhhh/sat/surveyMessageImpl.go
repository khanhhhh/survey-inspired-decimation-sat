package sat

type message = float64

func absMessage(messageIn message) (messageOut message) {
	if messageIn >= 0 {
		messageOut = +messageIn
	} else {
		messageOut = -messageIn
	}
	return messageOut
}
