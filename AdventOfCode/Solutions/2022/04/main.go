package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

// --- Day 4: Camp Cleanup ---
func run(input string) (interface{}, interface{}) {
	count, countb := 0, 0

	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, ",")
		left := helpers.Map(strings.Split(splits[0], "-"), pkg.MustAtoi)
		right := helpers.Map(strings.Split(splits[1], "-"), pkg.MustAtoi)

		if left[0] > right[0] {
			left, right = right, left
		}
		if left[1] >= right[0] {
			countb++

			// If left is strictly larger or if they start at the same
			// point and right is strictly larger
			if left[1] >= right[1] ||
				right[0] == left[0] && right[1] >= left[1] {
				count++
			}
		}
	}
	return count, countb
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
