package main

import "sat"

import "math/rand"

func main() {
	re := rand.New(rand.NewSource(1234154342))
	ins := sat.Random3SAT(re, 2048, 4.0)
	ins.Solve()
}
