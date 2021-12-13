package main

import (
	"strings"
	"unicode"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 12: Passage Pathing ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	caves := make(map[string][]string)

	for _, line := range lines {
		fromTo := strings.Split(line, "-")
		caves[fromTo[0]] = append(caves[fromTo[0]], fromTo[1])
		caves[fromTo[1]] = append(caves[fromTo[1]], fromTo[0])
	}
	numberOfRoutes := getRoutesLong(caves, make(map[string]bool), "start", -1)
	numberOfRoutes2 := getRoutesLong(caves, make(map[string]bool), "start", 0)

	return numberOfRoutes, numberOfRoutes2
}

func getRoutesLong(caves map[string][]string, seen map[string]bool, place string, bonusCave int) int {
	numberOfCaves := 0
	for _, placeToGo := range caves[place] {
		switch placeToGo {
		case "end":
			numberOfCaves += 1
			continue
		case "start":
			continue
		}

		if unicode.IsUpper(rune(placeToGo[0])) {
			numberOfCaves += getRoutesLong(caves, seen, placeToGo, bonusCave)
			continue
		}

		// Be efficient about memory allocation and map copies
		bonus := bonusCave
		newSeen := seen
		if seen[placeToGo] {
			if bonusCave == -1 || bonusCave == 1 {
				continue
			}
			bonus = 1
		} else {
			newSeen = map[string]bool{}
			for k, v := range seen {
				newSeen[k] = v
			}
			newSeen[placeToGo] = true
		}

		numberOfCaves += getRoutesLong(caves, newSeen, placeToGo, bonus)
	}

	return numberOfCaves
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
