package main

import (
	"image"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"golang.org/x/exp/maps"

	"adventOfCode/helpers/math"
)

// --- Day 14: Regolith Reservoir ---
func run(input string) (interface{}, interface{}) {
	maxY := 0
	cave := make(map[image.Point]int)

	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, " -> ")
		prevXY := math.Map(strings.Split(coords[0], ","), pkg.MustAtoi)
		if prevXY[1] > maxY {
			maxY = prevXY[1]
		}
		for i := 1; i < len(coords); i++ {
			currentXY := math.Map(strings.Split(coords[i], ","), pkg.MustAtoi)
			if currentXY[1] > maxY {
				maxY = currentXY[1]
			}
			if currentXY[0] > prevXY[0] {
				for x := prevXY[0]; x <= currentXY[0]; x++ {
					cave[image.Point{X: x, Y: prevXY[1]}] = 1
				}
			} else if currentXY[0] < prevXY[0] {
				for x := currentXY[0]; x <= prevXY[0]; x++ {
					cave[image.Point{X: x, Y: prevXY[1]}] = 1
				}
			} else if currentXY[1] > prevXY[1] {
				for y := prevXY[1]; y <= currentXY[1]; y++ {
					cave[image.Point{X: prevXY[0], Y: y}] = 1
				}
			} else if currentXY[1] < prevXY[1] {
				for y := currentXY[1]; y <= prevXY[1]; y++ {
					cave[image.Point{X: prevXY[0], Y: y}] = 1
				}
			}
			prevXY = currentXY
		}
	}
	drawCave(cave)
	copyMap := make(map[image.Point]int)
	maps.Copy(copyMap, cave)
	return runSand(cave, 500, 400), runSand(copyMap, 500, maxY+2)
}

func runSand(cave map[image.Point]int, x int, y int) int {
	source := image.Point{X: x, Y: 0}
	sand := 0
	for ; ; sand++ {
		currentSand := source

		for {
			if currentSand.Y+1 >= y {
				cave[currentSand] = 2
				break
			}

			down := image.Point{X: currentSand.X, Y: currentSand.Y + 1}
			left := image.Point{X: currentSand.X - 1, Y: currentSand.Y + 1}
			right := image.Point{X: currentSand.X + 1, Y: currentSand.Y + 1}
			if cave[down] != 1 && cave[down] != 2 {
				if down.Y > 300 {
					return sand
				}
				currentSand = down
			} else if cave[left] != 1 && cave[left] != 2 {
				currentSand = left
			} else if cave[right] != 1 && cave[right] != 2 {
				currentSand = right
			} else {
				if currentSand.Y == 0 {
					cave[currentSand] = 2
					drawCave(cave)
					return sand + 1
				}
				cave[currentSand] = 2
				break
			}
		}
	}
}

func drawCave(cave map[image.Point]int) {
	minX, minY, maxX, maxY := 420, 0, 620, 165

	for y := minY; y < maxY; y++ {
		line := ""
		for x := minX; x < maxX; x++ {
			switch cave[image.Point{X: x, Y: y}] {
			case 1:
				line += "#"
			case 2:
				line += "o"
			default:
				line += "."
			}
		}
		println(line)
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
