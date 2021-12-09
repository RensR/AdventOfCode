package main

import (
	"adventOfCode/helpers"
	"image"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	var lowests []image.Point

	lowest := int32(0)
	for y, line := range lines {
		for x, char := range line {
			// left + right
			if x > 0 && int32(line[x-1]) <= char || x < len(line)-1 && int32(line[x+1]) <= char {
				continue
			}
			// up + down
			if y > 0 && int32(lines[y-1][x]) <= char || y < len(lines)-1 && int32(lines[y+1][x]) <= char {
				continue
			}
			lowest += char - '0' + 1
			lowests = append(lowests, image.Point{
				X: x,
				Y: y,
			})
		}
	}

	bestCave, almostBest, stillGreat := 0, 0, 0
	found := make(map[image.Point]bool)
	for _, point := range lowests {
		newCave := checkAround(lines, point, found)
		if newCave >= bestCave {
			bestCave, almostBest, stillGreat = newCave, bestCave, almostBest
		} else if newCave >= almostBest {
			almostBest, stillGreat = newCave, almostBest
		} else if newCave > stillGreat {
			stillGreat = newCave
		}
	}

	return lowest, bestCave * almostBest * stillGreat
}

func checkAround(lines []string, point image.Point, found map[image.Point]bool) int {
	caveSize := 1
	found[point] = true

	for _, direction := range helpers.UpDownLeftRight {
		yy := point.Y + direction.Y
		xx := point.X + direction.X
		// Out of bounds check
		if yy < 0 || yy >= len(lines) || xx < 0 || xx >= len(lines[0]) {
			continue
		}
		newPoint := image.Point{X: xx, Y: yy}
		// Wall and already-checked check
		if lines[yy][xx] == '9' || found[newPoint] {
			continue
		}
		caveSize += checkAround(lines, newPoint, found)
	}

	return caveSize
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
