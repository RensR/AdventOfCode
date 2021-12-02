package main

import (
	"adventOfCode/helpers"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

type movement struct {
	vector twod.Vector
	distance int
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	exit, exit2,  err := parse(input, "\n")
	if err != nil {
		panic(err)
	}

	return helpers.Abs(int64(exit.Pos.X() * exit.Pos.Y())), helpers.Abs(int64(exit2.Pos.X() * exit2.Pos.Y()))
}

func parse(s string, sep string) (*twod.P, *twod.P, error){
	lines := strings.Split(s, sep)
	start := twod.NewPoint(0,0)
	start2 := twod.NewPoint(0,0)
	aim := 0
	for _, line := range lines {
		parts := strings.Split(line," ")
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		switch parts[0] {
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
	return start, start2, nil
}

func main() {
	execute.Run(run, tests, puzzle, true)
}