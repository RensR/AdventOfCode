package main

import (
	"image"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	lines := parse(input)
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
	y := line.from.Y

	for x := line.from.X; x <= line.to.X; x++ {
		makeMap(hitMap, x)
		hitMap[x][y]++
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
	x, _ := strconv.Atoi(xAndY[0])
	y, _ := strconv.Atoi(xAndY[1])
	return image.Point{
		X: x,
		Y: y,
	}
}

func parse(s string) (result []fromTo) {
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		parsedLine := fromTo{
			from: parsePoint(parts[0]),
			to:   parsePoint(parts[1]),
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
