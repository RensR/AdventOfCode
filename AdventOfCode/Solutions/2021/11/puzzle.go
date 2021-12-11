package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzle,
		`1688`,
		`403`,
	},
	{
		puzzleTest,
		`1656`,
		`195`,
	},
}

var puzzleTest = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var puzzle = `3172537688
4566483125
6374512653
8321148885
4342747758
1362188582
7582213132
6887875268
7635112787
7242787273`
