package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 21: Monkey Math ---
func run(input string) (interface{}, interface{}) {
	monkeys := GetMonkeyMap(input)
	return monkeys["root"].f(monkeys), YoureTheMonkeyNow(monkeys)
}

func GetMonkeyMap(input string) map[string]Math {
	monkeys := make(map[string]Math)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		math := strings.Split(parts[1], " ")
		if len(math) == 1 {
			monkeys[parts[0]] = Math{f: func(monkeys map[string]Math) int {
				return pkg.MustAtoi(math[0])
			}}
		} else {
			switch math[1] {
			case "+":
				monkeys[parts[0]] = Math{
					f: func(monkeys map[string]Math) int {
						return monkeys[math[0]].f(monkeys) + monkeys[math[2]].f(monkeys)
					},
					left:  math[0],
					right: math[2],
					op:    math[1],
				}
			case "-":
				monkeys[parts[0]] = Math{
					f: func(monkeys map[string]Math) int {
						return monkeys[math[0]].f(monkeys) - monkeys[math[2]].f(monkeys)
					},
					left:  math[0],
					right: math[2],
					op:    math[1],
				}
			case "/":
				monkeys[parts[0]] = Math{
					f: func(monkeys map[string]Math) int {
						return monkeys[math[0]].f(monkeys) / monkeys[math[2]].f(monkeys)
					},
					left:  math[0],
					right: math[2],
					op:    math[1],
				}
			case "*":
				monkeys[parts[0]] = Math{
					f: func(monkeys map[string]Math) int {
						return monkeys[math[0]].f(monkeys) * monkeys[math[2]].f(monkeys)
					},
					left:  math[0],
					right: math[2],
					op:    math[1],
				}
			}
		}
	}
	return monkeys
}

func YoureTheMonkeyNow(monkeys map[string]Math) int {
	lookingFor := 0
	currentNode := monkeys["root"]
	currentNode.op = "="

	for {
		monkeys["humn"] = Math{f: func(monkeys map[string]Math) int {
			return -1000
		}}

		leftVal := monkeys[currentNode.left].f(monkeys)
		rightVal := monkeys[currentNode.right].f(monkeys)

		monkeys["humn"] = Math{f: func(monkeys map[string]Math) int {
			return 1234567
		}}

		leftValNew := monkeys[currentNode.left].f(monkeys)

		if leftVal != leftValNew {
			// Left has changed, left contains humn
			// the value of _left_ `op` _right_ needs to be _lookingFor_
			// The new lookingFor should be left's value
			switch currentNode.op {
			case "+":
				lookingFor -= rightVal
			case "-":
				lookingFor += rightVal
			case "/":
				lookingFor *= rightVal
			case "*":
				lookingFor /= rightVal
			case "=":
				lookingFor = rightVal
			}
			if currentNode.left == "humn" {
				return lookingFor
			}
			currentNode = monkeys[currentNode.left]
		} else {
			switch currentNode.op {
			case "+":
				lookingFor -= leftVal
			case "-":
				lookingFor = leftVal - lookingFor
			case "/":
				lookingFor = leftVal / lookingFor
			case "*":
				lookingFor /= leftVal
			case "=":
				lookingFor = leftVal
			}
			if currentNode.right == "humn" {
				return lookingFor
			}
			currentNode = monkeys[currentNode.right]
		}
	}
}

type Math struct {
	f           func(monkeys map[string]Math) int
	left, right string
	op          string
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
