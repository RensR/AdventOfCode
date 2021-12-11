package main

import (
	"adventOfCode/helpers"
	"image"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 11: Dumbo Octopus ---
func run(input string) (interface{}, interface{}) {
	cave := helpers.ParseIntMap(input)

	totalFlashes := 0
	syncAfter := 0
	for i := 0; ; i++ {
		newCave, flashes := passTime(cave)
		cave = newCave
		if i < 100 {
			totalFlashes += flashes
		}
		if flashes == 100 {
			syncAfter = i + 1
			break
		}
	}

	return totalFlashes, syncAfter
}

func passTime(cave map[image.Point]int) (map[image.Point]int, int) {
	var flashes []image.Point

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			point := image.Point{X: x, Y: y}
			cave[point]++
			if cave[point] > 9 {
				flashes = append(flashes, point)
				cave[point] = -1000
			}
		}
	}
	flashCount := len(flashes)

	for len(flashes) > 0 {
		for _, direction := range helpers.AllDirections {
			flashedPoint := image.Point{X: flashes[0].X + direction.X, Y: flashes[0].Y + direction.Y}
			// Out of bounds check
			if flashedPoint.Y < 0 || flashedPoint.Y >= 10 || flashedPoint.X < 0 || flashedPoint.X >= 10 {
				continue
			}
			cave[flashedPoint]++
			if cave[flashedPoint] > 9 {
				flashCount++
				flashes = append(flashes, flashedPoint)
				cave[flashedPoint] = -1000
			}
		}
		flashes = flashes[1:]
	}

	// reset all flashed spots to 0
	for point, value := range cave {
		if value < 0 {
			cave[point] = 0
		}
	}
	return cave, flashCount
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
