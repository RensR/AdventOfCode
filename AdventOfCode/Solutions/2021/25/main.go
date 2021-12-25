package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 25: Sea Cucumber ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	ground := make(map[xy]int)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			switch lines[y][x] {
			case '.':
				continue
			case '>':
				ground[xy{x: x, y: y}] = 1
			case 'v':
				ground[xy{x: x, y: y}] = 2
			}
		}
	}

	max := xy{x: len(lines[0]), y: len(lines)}

	return move(ground, max), nil
}

func move(ground map[xy]int, max xy) int {
	hasMoved := true
	day := 0
	nextDay := make(map[xy]int)

	for hasMoved {
		day++
		hasMoved = false
		for loc, cucumber := range ground {
			if cucumber == 1 {
				next := (loc.x + 1) % max.x
				if _, ok := ground[xy{x: next, y: loc.y}]; ok {
					nextDay[loc] = cucumber
				} else {
					hasMoved = true
					nextDay[xy{x: next, y: loc.y}] = cucumber
				}
			} else {
				nextDay[loc] = cucumber
			}
		}
		ground = nextDay
		nextDay = make(map[xy]int)

		for loc, cucumber := range ground {
			if cucumber == 2 {
				next := (loc.y + 1) % max.y
				if _, ok := ground[xy{x: loc.x, y: next}]; ok {
					nextDay[loc] = cucumber
				} else {
					hasMoved = true
					nextDay[xy{x: loc.x, y: next}] = cucumber
				}
			} else {
				nextDay[loc] = cucumber
			}
		}
		ground = nextDay
		nextDay = make(map[xy]int)
	}

	return day
}

func main() {
	execute.Run(run, tests, puzzle, true)
}

type xy struct {
	x int
	y int
}
