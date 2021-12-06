package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (a1 interface{}, a2 interface{}) {
	ages := make([]int, 9)

	for _, num := range pkg.ParseIntList(input, ",") {
		ages[num]++
	}

	for i := 0; i < 256; i++ {
		if i == 80 {
			a1 = countFish(ages)
		}
		// shorter but slower answer
		// ages = []int{ages[1], ages[2], ages[3], ages[4], ages[5], ages[6], ages[7] + ages[0], ages[8], ages[0]}
		ages = passDay(ages)
	}

	return a1, countFish(ages)
}

func countFish(ages []int) (result int) {
	for _, age := range ages {
		result += age
	}
	return result
}

func passDay(ages []int) []int {
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
