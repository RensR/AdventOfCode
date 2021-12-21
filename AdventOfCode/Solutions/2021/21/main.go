package main

import (
	"adventOfCode/helpers"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 21: Dirac Dice ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	p1 := pkg.MustAtoi(strings.Split(lines[0], ": ")[1])
	p2 := pkg.MustAtoi(strings.Split(lines[1], ": ")[1])

	d100 := dice{value: 0, wraps: 0}
	p1Score, p2Score := 0, 0

	for p1Score < 1000 && p2Score < 1000 {
		p1 += d100.roll() + d100.roll() + d100.roll()
		p1Score += (p1-1)%10 + 1
		if p1Score >= 1000 {
			break
		}
		p2 += d100.roll() + d100.roll() + d100.roll()
		p2Score += (p2-1)%10 + 1
	}

	a1 := helpers.Min(p1Score, p2Score) * (d100.wraps*100 + d100.value)

	return a1, helpers.Max64(quant(
		int64(pkg.MustAtoi(strings.Split(lines[0], ": ")[1])),
		int64(pkg.MustAtoi(strings.Split(lines[1], ": ")[1]))),
	)
}

type board struct {
	p1      int64
	p2      int64
	p1Score int64
	p2Score int64
}

func quant(p1 int64, p2 int64) (int64, int64) {
	lookupP1 := make(map[board]board)
	lookupP2 := make(map[board]board)
	return quantumLeap(board{p1: p1, p2: p2, p1Score: 0, p2Score: 0}, lookupP1, lookupP2, true)
}

func quantumLeap(current board, lookupP1 map[board]board, lookupP2 map[board]board, player1 bool) (int64, int64) {
	if current.p1Score >= 21 {
		return 1, 0
	}
	if current.p2Score >= 21 {
		return 0, 1
	}

	newP1, newP2 := int64(0), int64(0)
	if player1 {
		if val, ok := lookupP1[current]; ok {
			return val.p1, val.p2
		}
		for _, throw := range getThrows() {
			b := board{p1: current.p1 + throw, p2: current.p2, p1Score: current.p1Score, p2Score: current.p2Score}
			b.p1Score += (b.p1-1)%10 + 1
			dp1, dp2 := quantumLeap(b, lookupP1, lookupP2, false)
			newP1 += dp1
			newP2 += dp2
		}
		lookupP1[current] = board{p1: newP1, p2: newP2}
	} else {
		if val, ok := lookupP2[current]; ok {
			return val.p1, val.p2
		}
		for _, throw := range getThrows() {
			b := board{p1: current.p1, p2: current.p2 + throw, p1Score: current.p1Score, p2Score: current.p2Score}
			b.p2Score += +(b.p2-1)%10 + 1
			dp1, dp2 := quantumLeap(b, lookupP1, lookupP2, true)
			newP1 += dp1
			newP2 += dp2
		}
		lookupP2[current] = board{p1: newP1, p2: newP2}
	}
	return newP1, newP2
}

func getThrows() (throws []int64) {
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				throws = append(throws, int64(a+b+c))
			}
		}
	}
	return throws
}

type dice struct {
	value int
	wraps int
}

func (d *dice) roll() int {
	if d.value >= 100 {
		d.value = 0
		d.wraps++
	}
	d.value++
	return d.value
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
