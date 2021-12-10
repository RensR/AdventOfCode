package main

import (
	"sort"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	score := 0
	var uncorrupted []int

	for _, line := range strings.Split(input, "\n") {
		currentLineScore := checkLine(line)
		score += currentLineScore
		if currentLineScore == 0 {
			uncorrupted = append(uncorrupted, fixLine(line))
		}
	}
	sort.Ints(uncorrupted)

	return score, uncorrupted[len(uncorrupted)/2]
}

func fixLine(line string) (score int) {
	var brackStack []rune

	for _, char := range line {
		switch char {
		case '(', '[', '<', '{':
			brackStack = append(brackStack, char)
		case ')', ']', '>', '}':
			brackStack = brackStack[:len(brackStack)-1]
		}
	}

	for len(brackStack) > 0 {
		n := len(brackStack) - 1
		score *= 5
		switch brackStack[n] {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
		brackStack = brackStack[:n]
	}

	return score
}

func checkLine(line string) int {
	var brackStack []rune

	for _, char := range line {
		n := len(brackStack) - 1
		switch char {
		case '(', '[', '<', '{':
			brackStack = append(brackStack, char)
		case ')':
			if brackStack[n] != '(' {
				return 3
			}
			brackStack = brackStack[:n]
		case ']':
			if brackStack[n] != '[' {
				return 57
			}
			brackStack = brackStack[:n]
		case '}':
			if brackStack[n] != '{' {
				return 1197
			}
			brackStack = brackStack[:n]
		case '>':
			if brackStack[n] != '<' {
				return 25137
			}
			brackStack = brackStack[:n]
		}
	}

	return 0
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
