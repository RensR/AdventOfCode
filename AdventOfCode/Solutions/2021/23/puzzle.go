package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzleTest,
		`12521`,
		`44169`,
	},
	{
		puzzle,
		`11516`,
		``,
	},
}

var test = `#############
#AA.......AD#
###.#B#C#.###
  #.#B#C#.#  
  #D#B#C#D#  
  #A#B#C#D#  
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
