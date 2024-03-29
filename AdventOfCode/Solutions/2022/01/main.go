package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers/math"
)

// --- Day 1: Calorie Counting ---
func run(input string) (interface{}, interface{}) {
	var allFood []int
	for _, elf := range strings.Split(input, "\n\n") {
		allFood = append(allFood, math.Sum(pkg.ParseIntList(elf, "\n")))
	}

	math.ReverseSort(allFood)

	return allFood[0], allFood[0] + allFood[1] + allFood[2]
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
