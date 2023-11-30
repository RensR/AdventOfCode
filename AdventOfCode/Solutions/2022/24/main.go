package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 24: Blizzard Basin ---
func run(input string) (interface{}, interface{}) {
	traversalCache := make(map[SnapShot]int)
	var mapCache []BlizzardMap

	return nil, nil
}

func TraverseBlizzard()

type SnapShot struct {
	X, Y, Time int
}

type BlizzardMap struct {
	Blizzards []Blizzard
	Map       [][]int
}

type Blizzard struct {
	X, Y   int
	Dx, Dy int
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
