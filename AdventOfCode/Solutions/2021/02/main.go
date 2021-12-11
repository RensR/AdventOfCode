package main

import (
	"adventOfCode/helpers"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// --- Day 2: Dive! ---
func run(input string) (interface{}, interface{}) {
	exit, exit2 := parse(strings.Split(input, "\n"))
	return helpers.Abs(int64(exit.Pos.X() * exit.Pos.Y())), helpers.Abs(int64(exit2.Pos.X() * exit2.Pos.Y()))
}

func parse(lines []string) (*twod.P, *twod.P) {
	start := twod.NewPoint(0, 0)
	start2 := twod.NewPoint(0, 0)
	aim := 0
	for _, line := range lines {
		directionAndDistance := strings.Split(line, " ")
		distance := pkg.MustAtoi(directionAndDistance[1])
		switch directionAndDistance[0] {
		case "forward":
			start.MoveRight(distance)
			start2.MoveRight(distance)
			start2.MoveUp(aim * distance)
			break
		case "up":
			start.MoveUp(distance)
			aim += distance
			break
		case "down":
			start.MoveDown(distance)
			aim -= distance
			break
		}
	}
	return start, start2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
