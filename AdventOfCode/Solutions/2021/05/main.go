package main

import (
	"image"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 5: Hydrothermal Venture ---
func run(input string) (interface{}, interface{}) {
	lines := parse(strings.Split(input, "\n"))
	return walkTheLine(lines, false), walkTheLine(lines, true)
}

type fromTo struct {
	from image.Point
	to   image.Point
}

func walkVertical(hitMap map[int]map[int]int, line fromTo) {
	makeMap(hitMap, line.from.X)
	x := hitMap[line.from.X]

	if line.from.Y <= line.to.Y {
		for y := line.from.Y; y <= line.to.Y; y++ {
			x[y]++
		}
	} else {
		for y := line.from.Y; y >= line.to.Y; y-- {
			x[y]++
		}
	}
}

func walkHorizontal(hitMap map[int]map[int]int, line fromTo) {
	for x := line.from.X; x <= line.to.X; x++ {
		makeMap(hitMap, x)
		hitMap[x][line.from.Y]++
	}
}

func walkDiagonal(hitMap map[int]map[int]int, line fromTo) {
	x := line.from.X
	y := line.from.Y
	for i := 0; line.from.X+i <= line.to.X; i++ {
		makeMap(hitMap, x+i)
		if line.from.Y <= line.to.Y {
			hitMap[x+i][y+i]++
		} else {
			hitMap[x+i][y-i]++
		}
	}
}

func walkTheLine(lines []fromTo, diagonal bool) (count int) {
	hitMap := make(map[int]map[int]int)
	for _, line := range lines {
		if line.from.X == line.to.X {
			walkVertical(hitMap, line)
		} else if line.from.Y == line.to.Y {
			walkHorizontal(hitMap, line)
		} else if diagonal {
			walkDiagonal(hitMap, line)
		}
	}

	for _, markerX := range hitMap {
		for _, markerY := range markerX {
			if markerY >= 2 {
				count++
			}
		}
	}
	return count
}

func makeMap(hitMap map[int]map[int]int, x int) {
	if hitMap[x] == nil {
		hitMap[x] = make(map[int]int)
	}
}

func parsePoint(point string) image.Point {
	xAndY := strings.Split(point, ",")
	return image.Point{
		X: pkg.MustAtoi(xAndY[0]),
		Y: pkg.MustAtoi(xAndY[1]),
	}
}

func parse(lines []string) (result []fromTo) {
	for _, line := range lines {
		fromAndTo := strings.Split(line, " -> ")
		parsedLine := fromTo{
			from: parsePoint(fromAndTo[0]),
			to:   parsePoint(fromAndTo[1]),
		}
		// Swap to always start a line at a lower x coordinate
		if parsedLine.from.X > parsedLine.to.X {
			parsedLine = fromTo{
				from: parsedLine.to,
				to:   parsedLine.from,
			}
		}
		result = append(result, parsedLine)
	}

	return result
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
