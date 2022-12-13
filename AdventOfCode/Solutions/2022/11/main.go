package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

type Monkey struct {
	Items  helpers.Queue[int]
	Op     func(int) int
	Test   func(int) int
	Throws int
}

// --- Day 11: Monkey in the Middle ---
func run(input string) (interface{}, interface{}) {
	topLevelDiv := 1
	var monkeysA, monkeysB []Monkey
	for _, monkDef := range strings.Split(input, "\n\n") {
		monkey, div := MakeMonk(monkDef)
		monkeysA = append(monkeysA, monkey)
		monkeysB = append(monkeysB, monkey)
		topLevelDiv *= div
	}

	return ThrowStuff(monkeysA, 20, func(i int) int { return i / 3 }),
		ThrowStuff(monkeysB, 10000, func(i int) int { return i % topLevelDiv })
}

func ThrowStuff(monkeys []Monkey, rounds int, div func(int) int) int {
	for r := 0; r < rounds; r++ {
		for i, _ := range monkeys {
			if monkeys[i].Items.IsEmpty() {
				continue
			}
			for !monkeys[i].Items.IsEmpty() {
				newItem := div(monkeys[i].Op(*monkeys[i].Items.Pop()))
				nextMonk := monkeys[i].Test(newItem)
				monkeys[nextMonk].Items.Push(newItem)
				monkeys[i].Throws++
			}
		}
	}

	high, higher := 0, 0

	for _, monkey := range monkeys {
		if monkey.Throws > high {
			if monkey.Throws >= higher {
				high, higher = higher, monkey.Throws
			} else {
				high = monkey.Throws
			}
		}
	}

	return high * higher
}

func MakeMonk(input string) (monk Monkey, div int) {
	lines := strings.Split(input, "\n")
	op := strings.Split(lines[2][23:], " ")
	queue := helpers.Queue[int]{Items: []int{}}
	for _, item := range helpers.Map(strings.Split(lines[1][18:], ", "), pkg.MustAtoi) {
		queue.Push(item)
	}

	div = pkg.MustAtoi(lines[3][21:])
	return Monkey{
		Items: queue,
		Op: func(i int) int {
			val := i
			if op[1] != "old" {
				val = pkg.MustAtoi(op[1])
			}
			switch op[0] {
			case "+":
				return i + val
			case "*":
				return i * val
			default:
				panic("AAHHHH")
			}
		},
		Test: func(i int) int {
			if i%div == 0 {
				return pkg.MustAtoi(lines[4][29:])
			}
			return pkg.MustAtoi(lines[5][30:])
		},
		Throws: 0,
	}, div
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
