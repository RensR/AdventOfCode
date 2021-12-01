package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := pkg.ParseIntList(input, "\n")
	prev := -1
	part1 := -1
	part2 := 0
	for i, measure := range list {
		if measure > prev {
			part1 ++
		}
		prev = measure

		if i >= 3 {
			if list[i - 2] + list[i-1] + list[i] > list[i - 3] + list[i-2] + list[i-1]{
				part2++
			}
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}