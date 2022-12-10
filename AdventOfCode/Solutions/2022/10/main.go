package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 10: Cathode-Ray Tube ---
func run(input string) (interface{}, interface{}) {
	println("")

	lines := strings.Split(input, "\n")
	totalSignalStrength := 0
	signalStrength := 1
	drawLine := ""
	inputIndex := 0
	busy := false
	for cycle := 1; inputIndex < len(lines); cycle++ {
		if (cycle+20)%40 == 0 {
			totalSignalStrength += signalStrength * cycle
		}
		if cycle%40 == 1 {
			println(drawLine)
			drawLine = ""
		}
		pixel := (cycle - 1) % 40
		if signalStrength == pixel || signalStrength+1 == pixel || signalStrength-1 == pixel {
			drawLine += "#"
		} else {
			drawLine += " "
		}
		action := strings.Split(lines[inputIndex], " ")
		if busy {
			busy = false
			signalStrength += pkg.MustAtoi(action[1])
			inputIndex++
			continue
		}

		switch action[0] {
		case "addx":
			busy = true
		case "noop":
			inputIndex++
		}
	}
	println(drawLine)

	return totalSignalStrength, nil
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
