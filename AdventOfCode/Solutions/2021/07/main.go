package main

import (
	"math"

	"adventOfCode/helpers"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 7: The Treachery of Whales ---
func run(input string) (interface{}, interface{}) {
	positions := pkg.ParseIntList(input, ",")

	total := 0
	for _, num := range positions {
		total += num
	}
	avg := total / len(positions)

	var min = math.MaxInt
	for i := 0; ; i++ {
		current := moveCrabs(positions, i)
		if current >= min {
			break
		}
		min = current
	}

	var min2 = math.MaxInt
	seen := map[int]int{0: 0, 1: 1}
	for i := avg; ; i++ {
		current := moveCrabsExpensive(positions, i, seen)
		if current >= min2 {
			break
		}
		min2 = current
	}

	return min, min2
}

func moveCrabs(pos []int, newPos int) (movement int) {
	for _, num := range pos {
		movement += helpers.Abs(num - newPos)
	}
	return movement
}

func moveCrabsExpensive(pos []int, newPos int, seen map[int]int) (movement int) {
	for _, num := range pos {
		movement += getFromMap(seen, helpers.Abs(num-newPos))
	}
	return movement
}

func getFromMap(seen map[int]int, distance int) int {
	if val, ok := seen[distance]; ok {
		return val
	} else {
		result := getFromMap(seen, distance-1)
		seen[distance] = result + distance
		return seen[distance]
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
