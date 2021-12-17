package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 17: Trick Shot ---
func run(input string) (interface{}, interface{}) {
	xAndY := strings.Split(strings.Split(input, "area: x=")[1], ", y=")
	minMaxX := strings.Split(xAndY[0], "..")
	minMaxY := strings.Split(xAndY[1], "..")
	minX := pkg.MustAtoi(minMaxX[0])
	maxX := pkg.MustAtoi(minMaxX[1])
	minY := pkg.MustAtoi(minMaxY[0])
	maxY := pkg.MustAtoi(minMaxY[1])

	return shootManyTimes(minX, maxX, minY, maxY)
}

func shootManyTimes(minX, maxX, minY, maxY int) (int, int) {
	bestHeight := -1
	targetHit := 0

	for x := 1; x <= maxX; x++ {
		for y := minY; y < 200; y++ {
			newHeight := shoot(minX, maxX, minY, maxY, x, y)
			if newHeight > -1 {
				targetHit++
				if newHeight > bestHeight {
					bestHeight = newHeight
				}
			}
		}
	}
	return bestHeight, targetHit
}

func shoot(minX, maxX, minY, maxY, dx, dy int) int {
	x, y, height := 0, 0, 0

	for x <= maxX && y >= minY {
		if x >= minX && y <= maxY {
			return height
		}

		x += dx
		y += dy
		if y > height {
			height = y
		}
		if dx > 0 {
			dx--
		}

		dy--
	}

	return -1
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
