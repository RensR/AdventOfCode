package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers/datastructures"
	"adventOfCode/helpers/math"
)

type Monkey struct {
	Items  datastructures.Queue[int]
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
			for !monkeys[i].Items.IsEmpty() {
				newItem := div(monkeys[i].Op(*monkeys[i].Items.Pop()))
				nextMonk := monkeys[i].Test(newItem)
				monkeys[nextMonk].Items.Push(newItem)
				monkeys[i].Throws++
			}
		}
	}

	throws := math.Map(monkeys, func(m Monkey) int {
		return m.Throws
	})

	math.ReverseSort(throws)

	return throws[0] * throws[1]
}

func MakeMonk(input string) (monk Monkey, div int) {
	lines := strings.Split(input, "\n")
	op := strings.Split(lines[2][23:], " ")
	queue := datastructures.Queue[int]{Items: []int{}}
	for _, item := range math.Map(strings.Split(lines[1][18:], ", "), pkg.MustAtoi) {
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
