package helpers

import (
	"image"
	"strings"
)

type direction struct {
	X int
	Y int
}

var UpDownLeftRight = []direction{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

var AllDirections = []direction{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: -1, Y: -1},
	{X: 1, Y: 1},
	{X: -1, Y: 1},
	{X: 1, Y: -1},
}

func ParseIntMap(input string) map[image.Point]int {
	result := make(map[image.Point]int)
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			result[image.Point{X: x, Y: y}] = int(char - '0')
		}
	}
	return result
}
