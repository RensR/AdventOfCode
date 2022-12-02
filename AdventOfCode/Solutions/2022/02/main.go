package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 2: Rock Paper Scissors ---
func run(input string) (interface{}, interface{}) {
	score := 0

	for _, line := range strings.Split(input, "\n") {
		switch line[0] {
		case 'A':
			switch line[2] {
			case 'X':
				score += 1 + 3
			case 'Y':
				score += 2 + 6
			case 'Z':
				score += 3 + 0
			}
		case 'B':
			switch line[2] {
			case 'X':
				score += 1 + 0
			case 'Y':
				score += 2 + 3
			case 'Z':
				score += 3 + 6
			}
		case 'C':
			switch line[2] {
			case 'X':
				score += 1 + 6
			case 'Y':
				score += 2 + 0
			case 'Z':
				score += 3 + 3
			}
		}
	}

	score2 := 0
	for _, line := range strings.Split(input, "\n") {
		switch line[0] {
		case 'A':
			switch line[2] {
			case 'X':
				score2 += 3 + 0
			case 'Y':
				score2 += 1 + 3
			case 'Z':
				score2 += 2 + 6
			}
		case 'B':
			switch line[2] {
			case 'X':
				score2 += 1 + 0
			case 'Y':
				score2 += 2 + 3
			case 'Z':
				score2 += 3 + 6
			}
		case 'C':
			switch line[2] {
			case 'X':
				score2 += 2 + 0
			case 'Y':
				score2 += 3 + 3
			case 'Z':
				score2 += 1 + 6
			}
		}
	}
	return score, score2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
