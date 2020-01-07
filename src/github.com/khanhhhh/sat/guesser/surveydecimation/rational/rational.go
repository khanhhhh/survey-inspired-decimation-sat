package rational

type Rat float64

// FromInt :
func FromInt(a int, b int) (ratOut Rat) {
	ratOut = Rat(float64(a) / float64(b))
	return ratOut
}

// FromFloat :
func FromFloat(floatIn float64) (ratOut Rat) {
	ratOut = Rat(floatIn)
	return ratOut
}

// ToFloat :
func (ratIn Rat) ToFloat() (floatOut float64) {
	return float64(ratIn)
}

// Abs
func Abs(ratIn Rat) (ratOut Rat) {
	if ratIn >= 0 {
		ratOut = +ratIn
	} else {
		ratOut = -ratIn
	}
	return ratOut
}

func (rat Rat) Sign() (signOut int) {
	if rat > 0 {
		signOut = +1
	}
	if rat == 0 {
		signOut = 0
	}
	if rat < 0 {
		signOut = -1
	}
	return signOut
}

// Add
func Add(ratIn1 Rat, ratIn2 Rat) (ratOut Rat) {
	ratOut = ratIn1 + ratIn2
	return ratOut
}

// Sub
func Sub(ratIn1 Rat, ratIn2 Rat) (ratOut Rat) {
	ratOut = ratIn1 - ratIn2
	return ratOut
}

// Mul
func Mul(ratIn1 Rat, ratIn2 Rat) (ratOut Rat) {
	ratOut = ratIn1 * ratIn2
	return ratOut
}

// Div
func Div(ratIn1 Rat, ratIn2 Rat) (ratOut Rat) {
	ratOut = ratIn1 / ratIn2
	return ratOut
}
