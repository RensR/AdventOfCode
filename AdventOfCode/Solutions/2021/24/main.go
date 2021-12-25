package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 24: Arithmetic Logic Unit ---
func run(input string) (interface{}, interface{}) {
	actions := lineToAction(strings.Split(input, "\n"))
	return getModelId(actions, true), getModelId(actions, false)
}

type Monad struct {
	mem   map[uint8]*int64
	tape  []int
	index *int
}

type Action struct {
	f       string
	literal bool
	a       uint8
	b       int64
}

func getModelId(actions []Action, high bool) int64 {
	prev := 0
	zValues := map[int64]int64{0: 0}
	for i := 1; i < len(actions); i++ {
		if actions[i].f == "inp" {
			zValues = tryAllDigits(actions[prev:i], zValues, high)
			prev = i
		}
	}
	return tryAllDigits(actions[prev:], zValues, high)[0]
}

func tryAllDigits(actions []Action, zVals map[int64]int64, high bool) map[int64]int64 {
	newZValues := make(map[int64]int64)
	isReduceStep := actions[4].b == 26
	for i := int64(1); i <= 9; i++ {
		for oldZ, modelId := range zVals {
			monad := NewMonad([]int{int(i)}, oldZ)
			monad.run(actions)
			if isReduceStep && *monad.mem['x'] != 0 {
				continue
			}
			if existing, ok := newZValues[*monad.mem['z']]; ok {
				if high && existing > modelId*int64(10)+i || !high && existing < modelId*int64(10)+i {
					continue
				}
			}
			newZValues[*monad.mem['z']] = modelId*10 + i
		}
	}
	return newZValues
}

func lineToAction(lines []string) (actions []Action) {
	for _, line := range lines {
		segments := strings.Split(line+" ?", " ")
		number, err := strconv.Atoi(segments[2])
		b := int64(segments[2][0])
		if err == nil {
			b = int64(number)
		}
		actions = append(actions, Action{
			f:       segments[0],
			literal: err == nil || segments[0] == "inp",
			a:       segments[1][0],
			b:       b,
		})
	}
	return actions
}

func (monad Monad) run(actions []Action) int64 {
	for _, a := range actions {
		value := a.b
		if !a.literal {
			value = *monad.mem[uint8(a.b)]
		}
		switch a.f {
		case "inp":
			monad.inp(a.a)
		case "add":
			monad.add(a.a, value)
		case "mul":
			monad.mul(a.a, value)
		case "div":
			monad.div(a.a, value)
		case "mod":
			monad.mod(a.a, value)
		case "eql":
			monad.eql(a.a, value)
		}
	}
	return *monad.mem['z']
}

func NewMonad(tape []int, oldZ int64) Monad {
	monad := Monad{mem: map[uint8]*int64{}}
	x, y, w, z, index := int64(0), int64(0), int64(0), oldZ, 0
	monad.mem['x'] = &x
	monad.mem['y'] = &y
	monad.mem['z'] = &z
	monad.mem['w'] = &w
	monad.index = &index
	monad.tape = tape
	return monad
}

func (monad Monad) inp(to uint8) {
	*monad.mem[to] = int64(monad.tape[*monad.index])
	*monad.index++
}

func (monad Monad) add(to uint8, from int64) {
	*monad.mem[to] += from
}

func (monad Monad) mul(to uint8, from int64) {
	*monad.mem[to] *= from
}

func (monad Monad) div(to uint8, from int64) {
	*monad.mem[to] /= from
}

func (monad Monad) mod(to uint8, from int64) {
	*monad.mem[to] = *monad.mem[to] % from
}

func (monad Monad) eql(to uint8, from int64) {
	if *monad.mem[to] == from {
		*monad.mem[to] = 1
	} else {
		*monad.mem[to] = 0
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
