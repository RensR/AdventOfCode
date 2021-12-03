package main

import (
	"adventOfCode/helpers"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	one, two, err := parse(input, "\n")
	if err != nil {
		panic(err)
	}

	return one, two
}

func parse(s string, sep string) (int64, int64, error) {
	lines := strings.Split(s, sep)
	occ := findOcc(lines)

	a, b := "", ""
	oxy, co2 := check(0, lines, occ, 0)

	for _, bit := range occ {
		if bit > len(lines)/2 {
			a += "1"
			b += "0"
		} else {
			a += "0"
			b += "1"
		}
	}

	return helpers.FromBitString(a) * helpers.FromBitString(b), helpers.FromBitString(oxy[0]) * helpers.FromBitString(co2[0]), nil
}

func findOcc(lines []string) (occ []int) {
	for _, line := range lines {
		for i, c := range line {
			if i >= len(occ) {
				occ = append(occ, 0)
			}
			if c == 49 {
				occ[i]++
			}
		}
	}
	return occ
}

func check(index int64, lines []string, occ []int, instrument int) (oxy []string, co2 []string) {
	for _, line := range lines {
		lookingForOne := occ[index]*2 >= len(lines)
		if lookingForOne && line[index] == 49 || !lookingForOne && line[index] == 48 {
			oxy = append(oxy, line)
		} else {
			co2 = append(co2, line)
		}
	}

	if len(oxy) > 1 && instrument > -1 {
		oxy, _ = check(index+1, oxy, findOcc(oxy), 1)
	}
	if len(co2) > 1 && instrument < 1 {
		_, co2 = check(index+1, co2, findOcc(co2), -1)
	}

	return oxy, co2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
