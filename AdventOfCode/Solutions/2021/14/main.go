package main

import (
	"math"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 14: Extended Polymerization ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	rules := make(map[string]string)

	for i := 2; i < len(lines); i++ {
		fromTo := strings.Split(lines[i], " -> ")
		rules[fromTo[0]] = fromTo[1]
	}

	lookup := make(map[Tuple]map[uint8]int)
	return countWithinPolymer(startMutation(rules, lines[0], 10, make(map[uint8]int), lookup)),
		countWithinPolymer(startMutation(rules, lines[0], 40, make(map[uint8]int), lookup))
}

func countWithinPolymer(counts map[uint8]int) int {
	minCount, maxCount := math.MaxInt, 0

	for _, value := range counts {
		if value > maxCount {
			maxCount = value
		}
		if value < minCount {
			minCount = value
		}
	}

	return maxCount - minCount
}

type Tuple struct {
	Depth  int
	Target string
}

func startMutation(rules map[string]string, polymer string, depth int, count map[uint8]int, lookup map[Tuple]map[uint8]int) map[uint8]int {
	count[polymer[0]]++
	for i := 1; i < len(polymer); i++ {
		result := mutateDeep(rules, string(polymer[i-1])+string(polymer[i]), depth-1, lookup)
		addMaps(count, result)
		count[polymer[i]]++
	}
	return count
}

func mutateDeep(rules map[string]string, polymer string, depth int, lookup map[Tuple]map[uint8]int) map[uint8]int {
	if val, ok := lookup[Tuple{Target: polymer, Depth: depth}]; ok {
		return val
	}

	result := make(map[uint8]int)
	if depth < 1 {
		if val, ok := rules[polymer]; ok {
			result[val[0]]++
		}
		return result
	}

	if val, ok := rules[polymer]; ok {
		addMaps(result, mutateDeep(rules, string(polymer[0])+val, depth-1, lookup))
		result[val[0]]++
		addMaps(result, mutateDeep(rules, val+string(polymer[1]), depth-1, lookup))
	}

	lookup[Tuple{Target: polymer, Depth: depth}] = result

	return result
}

func addMaps(a map[uint8]int, b map[uint8]int) {
	for k, v := range b {
		a[k] += v
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
