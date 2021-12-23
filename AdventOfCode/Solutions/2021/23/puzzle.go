package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		test,
		`16059`,
		``,
	},
	{
		puzzleTest,
		`12521`,
		`44169`,
	},
	{
		puzzle,
		`11516`,
		`40272`,
	},
}

var test = `#############
#...........#
###D#D#C#C###
  #B#A#B#A#  
  #########  `

var puzzleTest = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#  
  #########  `

var puzzle = `#############
#...........#
###C#A#B#D###
  #B#A#D#C#  
  #########  `
