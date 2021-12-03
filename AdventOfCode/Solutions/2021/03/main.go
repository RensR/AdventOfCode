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

	g, e := 0, 0
	oxy, co2 := check(0, lines, occ, 0)

	for _, bit := range occ {
		g <<= 1
		e <<= 1

		if bit > len(lines)/2 {
			g |= 1
		} else {
			e |= 1
		}
	}

	return int64(g * e), helpers.FromBitString(oxy[0]) * helpers.FromBitString(co2[0]), nil
}

func findOcc(lines []string) []int {
	occ := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				occ[i]++
			}
		}
	}
	return occ
}

func check(index int64, lines []string, occ []int, instrument int) (oxy []string, co2 []string) {
	for _, line := range lines {
		lookingForOne := occ[index]*2 >= len(lines)
		if lookingForOne && line[index] == '1' || !lookingForOne && line[index] == '0' {
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
