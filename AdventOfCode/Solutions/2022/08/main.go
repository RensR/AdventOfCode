package main

import (
	"image"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers/grid"
	_map "adventOfCode/helpers/map"
)

// --- Day 8: Treetop Tree House ---
func run(input string) (interface{}, interface{}) {
	treeLines := strings.Split(input, "\n")
	treeMap := grid.ParseIntMap(input)
	treeVisibilityMap := make(map[image.Point]int)

	height, width := len(treeLines), len(treeLines[0])

	for x := 0; x < width; x++ {
		maxHeight := -1
		for y := 0; y < height; y++ {
			if treeMap[image.Point{X: x, Y: y}] > maxHeight {
				treeVisibilityMap[image.Point{X: x, Y: y}]++
				maxHeight = treeMap[image.Point{X: x, Y: y}]
			}
		}
		maxHeight = -1
		for y := height - 1; y >= 0; y-- {
			if treeMap[image.Point{X: x, Y: y}] > maxHeight {
				treeVisibilityMap[image.Point{X: x, Y: y}]++
				maxHeight = treeMap[image.Point{X: x, Y: y}]
			}
		}
	}

	for y := 0; y < width; y++ {
		maxHeight := -1
		for x := 0; x < height; x++ {
			if treeMap[image.Point{X: x, Y: y}] > maxHeight {
				treeVisibilityMap[image.Point{X: x, Y: y}]++
				maxHeight = treeMap[image.Point{X: x, Y: y}]
			}
		}
		maxHeight = -1
		for x := height - 1; x >= 0; x-- {
			if treeMap[image.Point{X: x, Y: y}] > maxHeight {
				treeVisibilityMap[image.Point{X: x, Y: y}]++
				maxHeight = treeMap[image.Point{X: x, Y: y}]
			}
		}
	}

	treeScenicMap := make(map[image.Point]int)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			scenicScore := 1
			for _, direction := range grid.UpDownLeftRight {
				for i := 1; ; i++ {
					xs, ys := x+direction.X*i, y+direction.Y*i
					if xs < 0 || xs >= width || ys < 0 || ys >= height {
						scenicScore *= i - 1
						break
					}
					if treeMap[image.Point{X: xs, Y: ys}] >= treeMap[image.Point{X: x, Y: y}] {
						scenicScore *= i
						break
					}
				}
			}
			treeScenicMap[image.Point{X: x, Y: y}] = scenicScore
		}
	}

	return len(treeVisibilityMap), _map.MapMaxValue(treeScenicMap)
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
