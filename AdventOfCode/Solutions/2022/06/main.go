package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 6: Tuning Trouble ---
func run(input string) (interface{}, interface{}) {
	return calculate(input, 4), calculate(input, 14)
}

func calculate(input string, depth int) int {
	// Solution B
	dup := false
	for i := depth; i < len(input); i++ {
		dup = false
		for lookBack := 1; lookBack < depth && !dup; lookBack++ {
			for compareWith := 0; compareWith < lookBack; compareWith++ {
				if input[i-compareWith] == input[i-lookBack] {
					dup = true
					break
				}
			}
		}
		if !dup {
			return i + 1
		}
	}

	return -1
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
