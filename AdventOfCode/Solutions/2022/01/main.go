package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 1: Calorie Counting ---
func run(input string) (interface{}, interface{}) {
	max, max1, max2 := 0, 0, 0

	elves := strings.Split(input, "\n\n")

	for _, elf := range elves {
		food := pkg.ParseIntList(elf, "\n")
		totalFood := 0
		for _, f := range food {
			totalFood += f
		}

		if totalFood >= max {
			max2 = max1
			max1 = max
			max = totalFood
		} else if totalFood >= max1 {
			max2 = max1
			max1 = totalFood
		} else if totalFood >= max2 {
			max2 = totalFood
		}
	}

	return max, max + max1 + max2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
