package main

import (
	"adventOfCode/helpers"
	"image"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 13: Transparent Origami ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	dots := make(map[int]map[int]bool)
	var folds []image.Point
	atFolds := false
	maxY := 0

	for _, line := range lines {
		if line == "" {
			atFolds = true
			continue
		}
		if atFolds {
			text := strings.Split(line, "along ")
			directionAndAmount := strings.Split(text[1], "=")
			if directionAndAmount[0] == "x" {
				folds = append(folds, image.Point{
					X: pkg.MustAtoi(directionAndAmount[1]),
				})
			} else {
				folds = append(folds, image.Point{
					Y: pkg.MustAtoi(directionAndAmount[1]),
				})
			}
			continue
		}
		xAndY := strings.Split(line, ",")
		x := pkg.MustAtoi(xAndY[0])
		y := pkg.MustAtoi(xAndY[1])
		if dots[y] == nil {
			dots[y] = make(map[int]bool)
		}
		dots[y][x] = true
		maxY = helpers.Max(maxY, y)
	}

	for y := 0; y < maxY; y++ {
		if dots[y] == nil {
			dots[y] = make(map[int]bool)
		}
	}

	dots = fold(dots, folds[0])
	a1 := count(dots)

	for i := 1; i < len(folds); i++ {
		dots = fold(dots, folds[i])
	}

	return a1, printDots(dots)
}

func printDots(dots map[int]map[int]bool) string {
	for y := 0; y < 6; y++ {
		line := ""
		for x := 0; x < 40; x++ {
			if dots[y][x] {
				line += "X"
			} else {
				line += "."
			}
		}
		println(line)
	}
	println("DONE")
	return "CJHAZHKU"
}

func fold(dots map[int]map[int]bool, fold image.Point) map[int]map[int]bool {
	if fold.X > 0 {
		for y, yRange := range dots {
			for x, _ := range yRange {
				if x < fold.X {
					continue
				} else {
					dots[y][fold.X-(x-fold.X)] = true
					delete(yRange, x)
				}
			}
		}
	} else {
		oldLen := len(dots)
		for d := 1; fold.Y-d >= 0 && fold.Y+d < oldLen; d++ {
			dY := fold.Y + d
			for x, _ := range dots[dY] {
				dots[fold.Y-d][x] = true
			}
			delete(dots, fold.Y+d)
		}
	}
	return dots
}

func count(dots map[int]map[int]bool) int {
	result := 0
	for _, y := range dots {
		for _, x := range y {
			if x {
				result++
			}
		}
	}
	return result
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
