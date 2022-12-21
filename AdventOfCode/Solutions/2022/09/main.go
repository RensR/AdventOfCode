package main

import (
	"image"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers/math"
)

// --- Day 9: Rope Bridge ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	return simRope(lines, 2), simRope(lines, 10)
}

func simRope(input []string, ropeLength int) int {
	visited := map[image.Point]int{}
	var knots []image.Point

	for i := 0; i < ropeLength; i++ {
		knots = append(knots, image.Point{X: 0, Y: 0})
	}

	for _, line := range input {
		commands := strings.Split(line, " ")
		times := pkg.MustAtoi(commands[1])

		for i := 0; i < times; i++ {
			switch commands[0] {
			case "U":
				knots[0].Y++
			case "D":
				knots[0].Y--
			case "L":
				knots[0].X--
			case "R":
				knots[0].X++
			}
			updateRope(knots, visited)
		}
	}
	return len(visited)
}

func updateRope(knots []image.Point, visited map[image.Point]int) {
	for i := 1; i < len(knots); i++ {
		deltaX := knots[i-1].X - knots[i].X
		deltaY := knots[i-1].Y - knots[i].Y

		if math.Abs(deltaX) >= 2 {
			knots[i].X += (deltaX) / math.Abs(deltaX)
			if deltaY != 0 {
				knots[i].Y += (deltaY) / math.Abs(deltaY)
			}
		} else if math.Abs(deltaY) >= 2 {
			knots[i].Y += (deltaY) / math.Abs(deltaY)
			if deltaX != 0 {
				knots[i].X += (deltaX) / math.Abs(deltaX)
			}
		}
	}

	visited[knots[len(knots)-1]]++
}

func printKnots(knots []image.Point) {
	view := map[image.Point]int{}
	size := 18

	for i, knot := range knots {
		view[knot] = i
	}

	for y := size; y > -size; y-- {
		buffer := ""
		for x := -size; x < size; x++ {
			if val, ok := view[image.Point{X: x, Y: y}]; ok {
				buffer += strconv.FormatInt(int64(val), 10)
			} else {
				buffer += "."
			}
		}
		println(buffer + strconv.FormatInt(int64(y), 10))
	}
	println(" ")
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
