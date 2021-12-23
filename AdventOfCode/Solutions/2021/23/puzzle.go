package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzleTest,
		`12521`,
		``,
	},
	{
		puzzle,
		`11516`,
		``,
	},
}

var test = `#############
#.....D.....#
###.#B#C#D###
  #A#B#C#A#  
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
