package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

// --- Day 3: Rucksack Reorganization ---
func run(input string) (interface{}, interface{}) {
	priority, priority2 := 0, 0

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		left, right := helpers.SortString(line[len(line)/2:]), helpers.SortString(line[:len(line)/2])

		l, r := 0, 0
		for left[l] != right[r] {
			if left[l] < right[r] {
				l++
			} else {
				r++
			}
		}
		priority += getRuneValue(left[l])

		if (i+1)%3 == 0 {
			one, two, three := helpers.SortString(lines[i-2]), helpers.SortString(lines[i-1]), helpers.SortString(lines[i])
			o, tw, th := 0, 0, 0
			for one[o] != two[tw] || one[o] != three[th] {
				if one[o] < two[tw] && one[o] < three[th] {
					o++
				} else if two[tw] < three[th] {
					tw++
				} else {
					th++
				}
			}
			priority2 += getRuneValue(one[o])
		}

	}
	return priority, priority2
}

func getRuneValue(r uint8) int {
	if r < 97 {
		value := int(r) - int('A')
		return value + 1 + 26
	}
	value := int(r) - int('a')
	return value + 1
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
