package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

type location struct {
	X int
	Y int
}

var allDirections = []location{
	{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
	{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0},
	{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
}

// --- Day 20: Trench Map ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	gameOfLife := make(map[location]bool)

	for y := 2; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			gameOfLife[location{X: x, Y: y}] = lines[y][x] == '#'
		}
	}

	a1 := 0
	dark := false
	for day := 0; day < 50; day++ {
		if day == 2 {
			a1 = countResult(gameOfLife)
		}
		gameOfLife = live(gameOfLife, lines[0], dark)
		dark = !dark
	}

	return a1, countResult(gameOfLife)
}

func live(gameOfLife map[location]bool, lookup string, dark bool) map[location]bool {
	nextDay := make(map[location]bool)

	for oldPlace, _ := range gameOfLife {
		for _, direction := range allDirections {
			place := location{X: oldPlace.X + direction.X, Y: oldPlace.Y + direction.Y}
			if _, ok := nextDay[place]; ok {
				continue
			}
			index := 0
			for _, direct := range allDirections {
				index = index << 1
				if val, ok := gameOfLife[location{X: place.X + direct.X, Y: place.Y + direct.Y}]; ok && val {
					index |= 1
				} else if !ok && dark && lookup[0] == '#' {
					index |= 1
				}
			}
			nextDay[place] = lookup[index] == '#'
		}
	}

	return nextDay
}

func countResult(gameOfLife map[location]bool) (result int) {
	for _, val := range gameOfLife {
		if val {
			result++
		}
	}
	return result
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
