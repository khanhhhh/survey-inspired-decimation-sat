package main

import sat "sid-sat"

import "math/rand"

func main() {
	re := rand.New(rand.NewSource(1234154342))
	ins := sat.Random3SAT(re, 1000, 3.5)
	ins.Solve()
}
