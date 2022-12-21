package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers/datastructures"
)

// --- Day 5: Supply Stacks ---
func run(input string) (interface{}, interface{}) {
	parts := strings.Split(input, "\n\n")
	initString := strings.Split(parts[0], "\n")
	var stacksPartOne, stacksPartTwo = getInitialStack(initString), getInitialStack(initString)

	for _, line := range strings.Split(parts[1], "\n") {
		tempStack := datastructures.Stack[string]{}
		instructions := strings.Split(line, " ")
		amount, from, to := pkg.MustAtoi(instructions[1]), pkg.MustAtoi(instructions[3])-1, pkg.MustAtoi(instructions[5])-1
		for i := 0; i < amount; i++ {
			stacksPartOne[to].Push(stacksPartOne[from].Pop())
			tempStack.Push(stacksPartTwo[from].Pop())
		}
		for i := 0; i < amount; i++ {
			stacksPartTwo[to].Push(tempStack.Pop())
		}
	}

	var finalItemsA, finalItemsB string
	for i := 0; i < len(stacksPartOne); i++ {
		finalItemsA += stacksPartOne[i].Pop()
		finalItemsB += stacksPartTwo[i].Pop()
	}

	return finalItemsA, finalItemsB
}

func getInitialStack(initializationString []string) (stacks []datastructures.Stack[string]) {
	stackPart := strings.Split(initializationString[len(initializationString)-1], "   ")
	for _ = range stackPart {
		stacks = append(stacks, datastructures.Stack[string]{})
	}

	for i := len(initializationString) - 2; i >= 0; i-- {
		line := initializationString[i]
		for j := 0; j*4+1 < len(line); j++ {
			index := j*4 + 1
			if line[index] != ' ' {
				stacks[(index+4)/4-1].Push(string(line[index]))
			}
		}
	}
	return stacks
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
