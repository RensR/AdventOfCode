package main

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 13: Distress Signal ---
func run(input string) (interface{}, interface{}) {
	indexesOfRightOrder := 0

	var all [][]interface{}
	for i, line := range strings.Split(input, "\n\n") {
		leftRight := strings.Split(line, "\n")
		left := jsonParse(leftRight[0])
		right := jsonParse(leftRight[1])
		all = append(all, left, right)

		if comp(left, right) == 1 {
			indexesOfRightOrder += i + 1
		}
	}
	indexer1, indexer2 := jsonParse("[[2]]"), jsonParse("[[6]]")
	all = append(all, indexer1, indexer2)

	// Bubble sort
	for {
		swapped := false
		for i := 1; i < len(all); i++ {
			if comp(all[i-1], all[i]) == -1 {
				all[i-1], all[i] = all[i], all[i-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	decoderIndexes := 1
	for i, outerList := range all {
		if len(outerList) == 1 && reflect.TypeOf(outerList[0]).String() == list {
			innerList := outerList[0].([]interface{})
			if len(innerList) == 1 && reflect.TypeOf(innerList[0]).String() == number {
				number := innerList[0].(float64)
				if number == 6 || number == 2 {
					decoderIndexes *= i + 1
				}
			}
		}
	}

	return indexesOfRightOrder, decoderIndexes
}

const (
	number = "float64"
	list   = "[]interface {}"
)

func comp(left, right interface{}) int {
	// number & number case
	if reflect.TypeOf(left).String() == number && reflect.TypeOf(right).String() == number {
		leftNumber, rightNumber := left.(float64), right.(float64)
		if leftNumber < rightNumber {
			return 1
		}
		if leftNumber == rightNumber {
			return 0
		}
		return -1
	}
	// list and list case
	if reflect.TypeOf(left).String() == list && reflect.TypeOf(right).String() == list {
		leftList, rightList := left.([]interface{}), right.([]interface{})

		i := 0
		for ; i < len(leftList) && i < len(rightList); i++ {
			compResult := comp(leftList[i], rightList[i])
			if compResult != 0 {
				return compResult
			}
		}
		if len(leftList) == i && len(rightList) == i {
			return 0
		}
		if len(leftList) == i && len(rightList) > i {
			return 1
		}
		if len(leftList) > i && len(rightList) == i {
			return -1
		}

		return 0
	}
	if reflect.TypeOf(left).String() == list {
		return comp(left, []interface{}{right})
	}
	return comp([]interface{}{left}, right)
}

func jsonParse(input string) []interface{} {
	var result map[string][]interface{}
	if json.Unmarshal([]byte("{\"obj\":"+input+"}"), &result) != nil {
		panic("Json parsing error :o")
	}
	return result["obj"]
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
