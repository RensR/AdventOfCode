package main

import (
	"image"
	"math"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

// --- Day 12: Hill Climbing Algorithm ---
func run(input string) (interface{}, interface{}) {
	resultA, resultB := math.MaxInt, math.MaxInt
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == 'S' || char == 'a' {
				result := walk(input, helpers.Vertex{X: x, Y: y, Distance: 0})
				if result < resultB {
					resultB = result
				}
				if char == 'S' {
					resultA = result
				}
			}
		}
	}

	return resultA, resultB
}

func walk(input string, start helpers.Vertex) int {
	var q = helpers.NodeQueue{}
	q.NewQ()

	hills := make(map[image.Point]int32)

	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			hills[image.Point{X: x, Y: y}] = char
		}
	}
	visited := make(map[image.Point]int32)
	visited[image.Point{X: start.X, Y: start.Y}] = 0
	current := &start
	for ; ; current = q.Dequeue() {
		currentHeight := hills[image.Point{X: current.X, Y: current.Y}]
		if currentHeight == 'E' {
			return current.Distance
		}
		for _, direction := range helpers.UpDownLeftRight {
			next := image.Point{
				X: current.X + direction.X,
				Y: current.Y + direction.Y,
			}
			if location, ok := hills[next]; ok { // location within bounds
				if location <= currentHeight+1 || currentHeight == 'S' {
					if location == 'E' && currentHeight < 'y' {
						continue
					}
					if _, ok := visited[next]; !ok {
						q.Enqueue(helpers.Vertex{
							X:        next.X,
							Y:        next.Y,
							Distance: current.Distance + 1,
						})
						visited[next] = int32(current.Distance + 1)
					}
				}
			}
		}
		if q.IsEmpty() {
			return math.MaxInt
		}
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
