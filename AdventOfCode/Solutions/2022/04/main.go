package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers/math"
)

// --- Day 4: Camp Cleanup ---
func run(input string) (interface{}, interface{}) {
	countA, countB := 0, 0

	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, ",")
		left := math.Map(strings.Split(splits[0], "-"), pkg.MustAtoi)
		right := math.Map(strings.Split(splits[1], "-"), pkg.MustAtoi)

		if left[0] > right[0] {
			left, right = right, left
		}
		if left[1] >= right[0] {
			countB++

			// If left is strictly larger or if they start at the same
			// point and right is strictly larger
			if left[1] >= right[1] ||
				right[0] == left[0] && right[1] >= left[1] {
				countA++
			}
		}
	}
	return countA, countB
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
