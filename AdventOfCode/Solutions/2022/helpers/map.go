package helpers

import (
	"image"
	"strings"
)

type Location struct {
	X int
	Y int
}
type Direction struct {
	X int
	Y int
}

var UpDownLeftRight = []Direction{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

type XYZ struct {
	X, Y, Z int
}

func (xyz XYZ) Add(other XYZ) XYZ {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
	return xyz
}

var ZYXSides = []XYZ{
	{X: -1, Y: 0, Z: 0},
	{X: 1, Y: 0, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: 0, Z: -1},
	{X: 0, Y: 0, Z: 1},
}

var AllDirections = []Direction{
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
