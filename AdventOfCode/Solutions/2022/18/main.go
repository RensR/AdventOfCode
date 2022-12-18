package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

// --- Day 18: Boiling Boulders ---
func run(input string) (interface{}, interface{}) {
	clouds := make(map[helpers.XYZ]struct{})
	steam := make(map[helpers.XYZ]struct{})

	maxX, maxY, maxZ, min := 0, 0, 0, -1

	for _, particle := range strings.Split(input, "\n") {
		var x, y, z int
		_, err := fmt.Sscanf(particle, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic("")
		}
		maxX = helpers.Max(maxX, x)
		maxY = helpers.Max(maxY, y)
		maxZ = helpers.Max(maxZ, z)

		clouds[helpers.XYZ{X: x, Y: y, Z: z}] = struct{}{}
	}

	queue := helpers.Queue[helpers.XYZ]{}
	queue.Push(helpers.XYZ{})

	for !queue.IsEmpty() {
		current := *queue.Pop()
		for _, direction := range helpers.ZYXSides {
			side := current.Add(direction)
			if side.X < min || side.Y < min || side.Z < min ||
				side.X > maxX+1 || side.Y > maxY+1 || side.Z > maxZ+1 {
				continue
			}
			if _, ok := clouds[side]; !ok {
				if _, ok := steam[side]; !ok {
					steam[side] = struct{}{}
					queue.Push(side)
				}
			}
		}
	}

	sides, sidesB := 0, 0
	for particle, _ := range clouds {
		sidesClear, sidesBClear := 6, 6
		for _, direction := range helpers.ZYXSides {
			side := particle.Add(direction)
			if _, ok := clouds[side]; ok {
				sidesClear--
				sidesBClear--
			} else if _, ok := steam[side]; !ok {
				sidesBClear--
			}
		}
		sides += sidesClear
		sidesB += sidesBClear
	}

	return sides, sidesB
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
