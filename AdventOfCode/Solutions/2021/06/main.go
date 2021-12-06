package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	ages := parse(input)
	a1 := 0

	for i := 0; i < 256; i++ {
		if i == 80 {
			a1 = countFish(ages)
		}
		ages = passDay(ages)
	}

	return a1, countFish(ages)
}

func countFish(ages map[int]int) (result int) {
	for _, age := range ages {
		result += age
	}
	return result
}

func parse(s string) map[int]int {
	nums := pkg.ParseIntList(s, ",")
	ages := make(map[int]int)
	for i := 0; i < 9; i++ {
		ages[i] = 0
	}

	for _, num := range nums {
		ages[num]++
	}

	return ages
}

func passDay(ages map[int]int) map[int]int {
	dueFish := ages[0]
	for i := 1; i < 9; i++ {
		ages[i-1] = ages[i]
	}

	ages[8] = dueFish
	ages[6] += dueFish

	return ages
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
