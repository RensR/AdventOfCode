package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 1: Sonar Sweep ---
func run(input string) (interface{}, interface{}) {
	list := pkg.ParseIntList(input, "\n")
	prev, part1, part2 := -1, -1, 0
	for i, measure := range list {
		if measure > prev {
			part1++
		}
		prev = measure

		if i >= 3 && list[i] > list[i-3] {
			part2++
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
