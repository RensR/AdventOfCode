package main

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"golang.org/x/exp/slices"
)

// --- Day 13: Distress Signal ---
func run(input string) (interface{}, interface{}) {
	indexesOfRightOrder := 0
	all := [][]interface{}{jsonParse("[[2]]"), jsonParse("[[6]]")}

	for i, line := range strings.Split(input, "\n\n") {
		leftRight := strings.Split(line, "\n")
		left, right := jsonParse(leftRight[0]), jsonParse(leftRight[1])
		all = append(all, left, right)

		if comp(left, right) == 1 {
			indexesOfRightOrder += i + 1
		}
	}

	slices.SortFunc(all, func(l, r []interface{}) bool {
		return comp(l, r) == 1
	})

	decoderIndexes := 1
	for i, outerList := range all {
		if len(outerList) == 1 && isList(outerList[0]) {
			innerList := outerList[0].([]interface{})
			if len(innerList) == 1 && isNumber(innerList[0]) {
				number := innerList[0].(float64)
				if number == 6 || number == 2 {
					decoderIndexes *= i + 1
				}
			}
		}
	}

	return indexesOfRightOrder, decoderIndexes
}

func isNumber(i interface{}) bool {
	return reflect.TypeOf(i).String() == "float64"
}

func isList(i interface{}) bool {
	return reflect.TypeOf(i).String() == "[]interface {}"
}

func comp(left, right interface{}) int {
	// number & number case
	if isNumber(left) && isNumber(right) {
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
	if isList(left) && isList(right) {
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
	if isList(left) {
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
