package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzleTest,
		`739785`,
		`444356092776315`,
	},
	{
		puzzle,
		`1073709`,
		`148747830493442`,
	},
}

var puzzleTest = `Player 1 starting position: 4
Player 2 starting position: 8`

var puzzle = `Player 1 starting position: 9
Player 2 starting position: 3`
