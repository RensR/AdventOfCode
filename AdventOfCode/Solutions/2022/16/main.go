package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/RyanCarrier/dijkstra"
)

// --- Day 16: Proboscidea Volcanium ---
func run(input string) (interface{}, interface{}) {
	locations := make(map[string]Location)
	var unOpened []string

	nameToInt := make(map[string]int)

	for i, line := range strings.Split(input, "\n") {
		var name string
		var flow int
		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &name, &flow)
		if err != nil {
			panic(err)
		}
		
		nameToInt[name] = i

		trimmed := strings.TrimPrefix(strings.TrimPrefix(strings.Split(line, "to valve")[1], "s"), " ")

		if flow > 0 {
			unOpened = append(unOpened, name)
		}

		locations[name] = Location{
			FlowRate: flow,
			Tunnels:  strings.Split(trimmed, ", "),
		}
	}

	slices.Sort(unOpened)

	value, path := DigDigDig(make(map[string]int), locations, "AA", 30, 0, slices.Clone(unOpened))
	println(path)
	valueB, path := DigDigDigElly(make(map[string]int), locations, "AA", "AA", 26, 0, unOpened)
	println(path)

	return value, valueB
}

type Location struct {
	FlowRate int
	Tunnels  []string
}

func DigDigDigElly(cache map[string]int, locations map[string]Location, me string, elly string, time int, gas int, unOpened []string) (int, string) {
	if time <= 0 || len(unOpened) == 0 {
		return gas, " END"
	}
	if time < 5 && gas < 100 {
		return gas, " LOW GAS STOP"
	}
	// Check if we have been in this position but have released more gas
	snapshot := ""
	if me < elly {
		snapshot = SnapShot(me+elly, time, unOpened)
	} else {
		snapshot = SnapShot(elly+me, time, unOpened)
	}

	if val, ok := cache[snapshot]; ok {
		if gas > val {
			cache[snapshot] = gas
		} else {
			return val, " Path already explored"
		}
	}

	maxGasReleased := 0
	bestPath := ""

	if locations[elly].FlowRate > 0 && locations[me].FlowRate > 0 && locations[elly].FlowRate != locations[me].FlowRate {
		newUnOpened := UpdateUnOpened(unOpened, me)
		index := sort.SearchStrings(newUnOpened, elly)
		newUnOpened = slices.Delete(newUnOpened, index, index+1)

		newMap := maps.Clone(locations)
		newMap[elly] = Location{Tunnels: newMap[elly].Tunnels}
		newMap[me] = Location{Tunnels: newMap[me].Tunnels}

		totalGas := locations[elly].FlowRate*(time-1) + locations[me].FlowRate*(time-1)
		gas, _ := DigDigDigElly(cache, newMap, me, elly, time-1, gas, newUnOpened)

		totalGas += gas
		if totalGas > maxGasReleased {
			maxGasReleased = totalGas
			//bestPath = PrintStatusLine(time, me, elly, totalGas, path)
		}
	}

	// Try to release the current valve
	if locations[me].FlowRate > 0 {
		for _, tunnelElly := range locations[elly].Tunnels {
			newMap, totalGas := CloseValve(locations, me, time)
			newUnOpened := UpdateUnOpened(unOpened, me)

			gas, _ := DigDigDigElly(cache, newMap, me, tunnelElly, time-1, gas, newUnOpened)
			totalGas += gas
			if totalGas >= maxGasReleased {
				maxGasReleased = totalGas
				//bestPath = PrintStatusLine(time, me, tunnelElly, totalGas, path)
			}
		}
	}
	if locations[elly].FlowRate > 0 {
		for _, tunnel := range locations[me].Tunnels {
			newMap, totalGas := CloseValve(locations, elly, time)
			newUnOpened := UpdateUnOpened(unOpened, elly)

			gas, _ := DigDigDigElly(cache, newMap, tunnel, elly, time-1, gas, newUnOpened)
			totalGas += gas
			if totalGas > maxGasReleased {
				maxGasReleased = totalGas
				//bestPath = PrintStatusLine(time, tunnel, elly, totalGas, path)
			}
		}
	}

	if time > 1 {
		for _, tunnel := range locations[me].Tunnels {
			for _, tunnelElly := range locations[elly].Tunnels {
				totalGas, _ := DigDigDigElly(cache, maps.Clone(locations), tunnel, tunnelElly, time-1, gas, unOpened)
				if totalGas > maxGasReleased {
					maxGasReleased = totalGas
					//bestPath = PrintStatusLine(time, tunnel, tunnelElly, totalGas, path)
				}
			}
		}
	}

	cache[snapshot] = maxGasReleased
	return maxGasReleased, bestPath
}

func CloseValve(locations map[string]Location, valve string, time int) (map[string]Location, int) {
	newMap := maps.Clone(locations)
	newMap[valve] = Location{Tunnels: newMap[valve].Tunnels}

	return newMap, locations[valve].FlowRate * (time - 1)
}

func UpdateUnOpened(unOpened []string, valve string) []string {
	newUnOpened := slices.Clone(unOpened)
	index := sort.SearchStrings(newUnOpened, valve)
	return slices.Delete(newUnOpened, index, index+1)
}

func PrintStatusLine(time int, me, elly string, gas int, oldPath string) string {
	return fmt.Sprintf("%2d | me   | move %s   GAS %d "+
		"   | elly | move %s "+
		"\n%s", 27-time, me, gas, elly, oldPath)
}

func DigDigDig(cache map[string]int, locations map[string]Location, start string, time int, gas int, unOpened []string) (int, string) {
	if time <= 0 || len(unOpened) == 0 {
		return gas, " END"
	}

	// Check if we have been in this position but have released more gas
	snapshot := SnapShot(start, time, unOpened)
	if val, ok := cache[snapshot]; ok {
		if gas > val {
			cache[snapshot] = gas
		} else {
			return val, " Path already explored"
		}
	}

	maxGasReleased := 0
	bestPath := ""

	// Try to release the current valve
	if locations[start].FlowRate > 0 {
		newMap, totalGas := CloseValve(locations, start, time)

		newUnOpened := slices.Clone(unOpened)
		index := sort.SearchStrings(newUnOpened, start)
		newUnOpened = slices.Delete(newUnOpened, index, index+1)

		gas, path := DigDigDig(cache, newMap, start, time-1, gas, newUnOpened)
		totalGas += gas
		if totalGas >= maxGasReleased {
			maxGasReleased = totalGas
			bestPath = fmt.Sprintf("%2d | Open %s   GAS %d \n%s", 31-time, start, totalGas, path)
		}
	}

	for _, tunnel := range locations[start].Tunnels {
		newMap := make(map[string]Location)
		maps.Copy(newMap, locations)
		totalGas, path := DigDigDig(cache, newMap, tunnel, time-1, gas, unOpened)
		if totalGas > maxGasReleased {
			maxGasReleased = totalGas
			bestPath = fmt.Sprintf("%2d | GOTO %s   GAS %d \n%s", 31-time, tunnel, totalGas, path)
		}
	}

	cache[snapshot] = maxGasReleased

	return maxGasReleased, bestPath
}

func SnapShot(location string, timeLeft int, unOpened []string) string {
	return fmt.Sprintf("%s-%2d-%+v", location, timeLeft, unOpened)
}

func calcReachabilityMatrix(reachability map[string][]string, mappingInt map[string]int) map[string]map[string]int {
	graph := dijkstra.NewGraph()

	for k := range reachability {
		graph.AddVertex(mappingInt[k])
	}

	for k, v := range reachability {
		for _, l := range v {
			graph.AddArc(mappingInt[k], mappingInt[l], 1)
		}
	}

	matrix := map[string]map[string]int{}
	for k1, v1 := range mappingInt {
		matrix[k1] = map[string]int{}
		for k2, v2 := range mappingInt {
			best, _ := graph.Shortest(v1, v2)
			matrix[k1][k2] = int(best.Distance)
		}
	}

	return matrix
}

func main() {
	execute.Run(run, tests, puzzleTest, true)
}
