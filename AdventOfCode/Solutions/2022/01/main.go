package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

// --- Day 1: Calorie Counting ---
func run(input string) (interface{}, interface{}) {
	max, max1, max2 := 0, 0, 0

	for _, elf := range strings.Split(input, "\n\n") {
		food := helpers.Sum(pkg.ParseIntList(elf, "\n"))
		if food >= max {
			max2, max1, max = max1, max, food
		} else if food >= max1 {
			max2, max1 = max1, food
		} else if food >= max2 {
			max2 = food
		}
	}

	return max, max + max1 + max2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
