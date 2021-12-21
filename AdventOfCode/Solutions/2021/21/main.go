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
	P1 := int64(pkg.MustAtoi(strings.Split(lines[0], ": ")[1]))
	P2 := int64(pkg.MustAtoi(strings.Split(lines[1], ": ")[1]))

	a1 := leap(board{p1: P1, p2: P2, p1Score: 0, p2Score: 0, playerOneTurn: true})
	a2 := helpers.Max64(quantumLeap(board{p1: P1, p2: P2, p1Score: 0, p2Score: 0, playerOneTurn: true}, make(map[board]board)))
	return a1, a2
}

type board struct {
	p1            int64
	p2            int64
	p1Score       int
	p2Score       int
	playerOneTurn bool
}

type dice struct {
	value int64
	wraps int64
}

func (d *dice) roll() int64 {
	if d.value >= 100 {
		d.value = 0
		d.wraps++
	}
	d.value++
	return d.value
}

func (b *board) Move(steps int64) {
	if b.playerOneTurn {
		b.p1 += steps
		b.p1Score += int((b.p1-1)%10 + 1)
	} else {
		b.p2 += steps
		b.p2Score += int((b.p2-1)%10 + 1)
	}
	b.playerOneTurn = !b.playerOneTurn
}

func leap(b board) int64 {
	d100 := dice{value: 0, wraps: 0}

	for b.p1Score < 1000 && b.p2Score < 1000 {
		b.Move(d100.roll() + d100.roll() + d100.roll())
	}
	return helpers.Min64(int64(b.p1Score), int64(b.p2Score)) * (d100.wraps*100 + d100.value)
}

func quantumLeap(current board, lookup map[board]board) (p1Wins int64, p2Wins int64) {
	if current.p1Score >= 21 {
		return 1, 0
	} else if current.p2Score >= 21 {
		return 0, 1
	} else if val, ok := lookup[current]; ok {
		return val.p1, val.p2
	}

	for _, throw := range getThrows() {
		var b = board{p1: current.p1, p2: current.p2, p1Score: current.p1Score, p2Score: current.p2Score, playerOneTurn: current.playerOneTurn}
		b.Move(throw)
		dp1, dp2 := quantumLeap(b, lookup)
		p1Wins += dp1
		p2Wins += dp2
	}
	lookup[current] = board{p1: p1Wins, p2: p2Wins}
	return p1Wins, p2Wins
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

func main() {
	execute.Run(run, tests, puzzle, true)
}
