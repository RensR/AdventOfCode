package main

import (
	"adventOfCode/helpers"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

type sevenSegmentDisplay struct {
	segments []string
	values   []string
}

// --- Day 8: Seven Segment Search ---
func run(input string) (interface{}, interface{}) {
	var displays []sevenSegmentDisplay

	for _, line := range strings.Split(input, "\n") {
		digitsAndTarget := strings.Split(line, " | ")
		segments := strings.Split(digitsAndTarget[0], " ")
		val := strings.Split(digitsAndTarget[1], " ")
		// sort them to enable the lookup table in the future
		for i, s := range segments {
			segments[i] = helpers.SortString(s)
		}
		for i, s := range val {
			val[i] = helpers.SortString(s)
		}
		displays = append(displays, sevenSegmentDisplay{
			segments: segments,
			values:   val,
		})
	}

	totalA := 0
	for _, segment := range displays {
		totalA += solveBasicSegments(segment)
	}

	totalB := 0
	for _, segment := range displays {
		totalB += solveAllSegments(segment)
	}

	return totalA, totalB
}

// Frequency of each possible segment
// a = 8
// b = 6
// c = 8
// d = 7
// e = 4
// f = 9
// g = 7
func solveAllSegments(display sevenSegmentDisplay) (result int) {
	lookupTable := make(map[string]int)
	countPositions := make(map[rune]int)
	for _, displayPos := range display.segments {
		for _, char := range displayPos {
			countPositions[char]++
		}
	}

	positionSheet := make(map[int][]rune)
	for char, val := range countPositions {
		positionSheet[val] = append(positionSheet[val], char)
	}

	for _, displayPos := range display.segments {
		switch len(displayPos) {
		case 2:
			lookupTable[displayPos] = 1
		case 3:
			lookupTable[displayPos] = 7
		case 4:
			lookupTable[displayPos] = 4
		case 7:
			lookupTable[displayPos] = 8
		case 5: // 2, 3 or 5
			// Only 2 has an empty space bottom-right
			if !strings.ContainsRune(displayPos, positionSheet[9][0]) {
				lookupTable[displayPos] = 2
				continue
			}
			// if they have top and top-right it's 3 else 5
			if strings.ContainsRune(displayPos, positionSheet[8][0]) && strings.ContainsRune(displayPos, positionSheet[8][1]) {
				lookupTable[displayPos] = 3
				continue
			}
			lookupTable[displayPos] = 5
		case 6: // 0, 6 or 9
			// if they have bottom left they can be - 0 6 -
			if strings.ContainsRune(displayPos, positionSheet[4][0]) {
				// if they have top and top-right it's 0 else 6
				if strings.ContainsRune(displayPos, positionSheet[8][0]) && strings.ContainsRune(displayPos, positionSheet[8][1]) {
					lookupTable[displayPos] = 0
					continue
				}
				lookupTable[displayPos] = 6
				continue
			}
			lookupTable[displayPos] = 9
		}
	}

	var resultString = ""
	for _, value := range display.values {
		resultString += strconv.FormatInt(int64(lookupTable[value]), 10)
	}

	return pkg.MustAtoi(resultString)
}

func solveBasicSegments(display sevenSegmentDisplay) (result int) {
	for _, val := range display.values {
		switch len(val) {
		case 2, 3, 4, 7:
			result++
		}
	}
	return result
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
