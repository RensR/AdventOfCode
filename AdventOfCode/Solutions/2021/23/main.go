package main

import (
	"adventOfCode/helpers"
	"fmt"
	"strings"
	"sync"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var letters = []uint8{'A', 'B', 'C', 'D'}
var moveCost = map[int32]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

// --- Day 23: Amphipod ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	return runSimulation(parseInput(lines)).Distance, run2(lines)
}

func run2(lines []string) int {
	end := []string{lines[3], lines[4]}
	lines = append(append(lines[0:3], "  #D#C#B#A#  ", "  #D#B#A#C#  "), end...)
	return runSimulation(parseInput(lines)).Distance
}

func runSimulation(layout [][]uint8, pods []Amphipod) *Move {
	Q := NodeQ{}
	Q.NewQ()
	currentState := &Move{Distance: 0, state: State{pods, layout}}
	seenStates := make(map[string]int)
	for ; !currentState.state.IsFinal(); currentState = Q.Dequeue() {
		moves := currentState.state.getValidMoves(seenStates, currentState.Distance)
		for _, move := range moves {
			move.oldMove = currentState
			Q.Enqueue(move)
		}
	}
	printPathTaken(*currentState)
	return currentState
}

func printPathTaken(move Move) {
	printing := &move
	for printing != nil {
		(*printing).state.GetSignature(true)
		printing = (*printing).oldMove
	}
}

func parseInput(lines []string) ([][]uint8, []Amphipod) {
	var layout = make([][]uint8, len(lines[0]))
	var pods []Amphipod
	index := 0
	for y, line := range lines {
		for x, c := range line {
			layout[x] = append(layout[x], uint8(c))
			switch c {
			case 'A', 'B', 'C', 'D':
				pods = append(pods, Amphipod{
					index:    index,
					x:        x,
					y:        y,
					char:     uint8(c),
					moveCost: moveCost[c],
				})
				index++
			}
		}
	}
	return layout, pods
}

func (state State) GetSignature(print bool) string {
	totalString := ""
	for y := 1; y < len(state.layout[0])-1; y++ {
		line := ""
		for x := range state.layout {
			line += string(state.layout[x][y])
		}
		totalString += line + "\n"
	}
	if print {
		fmt.Println(totalString)
	}
	return totalString
}

func (state State) getValidMoves(lookup map[string]int, currentDistance int) (moves []Move) {
	for _, pod := range state.pods {
		if pod.IsDone(state) {
			continue
		}
		seenThisPod := map[helpers.Location]bool{helpers.Location{X: pod.x, Y: pod.y}: true}
		newMoves := DoStep(state, pod, 0, seenThisPod, pod.y < 2)
		for _, move := range newMoves {
			move.Distance += currentDistance
			signature := move.state.GetSignature(false)
			if move.state.IsFinal() {
				moves = append(moves, move)
				continue
			}

			if val, ok := lookup[signature]; ok && move.Distance >= val {
				continue
			}

			lookup[signature] = move.Distance
			moves = append(moves, move)
		}
	}
	return moves
}

func DoStep(state State, pod Amphipod, steps int, seenThisPod map[helpers.Location]bool, fromHallway bool) (moves []Move) {
	steps++
	for _, direction := range helpers.UpDownLeftRight {
		xx := pod.x + direction.X
		yy := pod.y + direction.Y
		if seenThisPod[helpers.Location{X: xx, Y: yy}] {
			continue
		}
		seenThisPod[helpers.Location{X: xx, Y: yy}] = true
		switch state.layout[xx][yy] {
		case '#', 'A', 'B', 'C', 'D':
			continue
		case '.':
			newState := makeCopyState(state, pod, xx, yy)
			switch xx {
			case 3, 5, 7, 9:
				switch direction.Y {
				case 0, -1: // going left/right/up
					moves = append(moves, DoStep(newState, newState.pods[pod.index], steps, seenThisPod, fromHallway)...)
				case 1: // going down
					// if he's going in, but it's not his room
					if xx != getXForPodId(pod.char) {
						continue
					}
					if newState.pods[pod.index].IsDone(newState) {
						moves = append(moves, Move{
							Distance: steps * pod.moveCost,
							state:    newState,
						})
					}
					wrongEntitiesInStack := false
					for y := len(state.layout[0]) - 2; y > 1 && !wrongEntitiesInStack; y-- {
						if state.layout[xx][y] != pod.char && state.layout[xx][y] != '.' {
							wrongEntitiesInStack = true
						}
					}
					if !wrongEntitiesInStack {
						moves = append(moves, DoStep(newState, newState.pods[pod.index], steps, seenThisPod, fromHallway)...)
					}
				}
			default:
				// add this move and do another step
				if !fromHallway {
					moves = append(moves, Move{
						Distance: steps * pod.moveCost,
						state:    newState,
					})
				}
				moves = append(moves, DoStep(newState, newState.pods[pod.index], steps, seenThisPod, fromHallway)...)
			}
		}
	}
	return moves
}

func makeCopyState(state State, pod Amphipod, xx int, yy int) State {
	newState := State{pods: []Amphipod{}, layout: makeCopyLayout(state.layout)}
	newState.layout[pod.x][pod.y] = '.'
	newState.layout[xx][yy] = pod.char
	for _, amphi := range state.pods {
		if amphi.index == pod.index {
			amphi.x = xx
			amphi.y = yy
		}
		newState.pods = append(newState.pods, amphi)
	}
	return newState
}

func makeCopyLayout(layout [][]uint8) (result [][]uint8) {
	cpy := make([][]uint8, len(layout))
	for i, val := range layout {
		deepCopy := make([]uint8, len(val))
		copy(deepCopy, val)
		cpy[i] = deepCopy
	}
	return cpy
}

// IsDone checks if a given Amphipod never needs to move again
func (pod Amphipod) IsDone(state State) bool {
	x := getXForPodId(pod.char)
	if pod.x != x {
		return false
	}

	for y := pod.y + 1; y < len(state.layout[0])-1; y++ {
		if state.layout[x][y] != pod.char {
			return false
		}
	}
	return true
}

func getXForPodId(podId uint8) int {
	return int((podId-'A')*2) + 3
}

// IsFinal returns true if all Amphipods are in the proper room
func (state State) IsFinal() bool {
	for _, letter := range letters {
		x := getXForPodId(letter)

		for y := 2; y < len(state.layout[0])-1; y++ {
			if state.layout[x][y] != letter {
				return false
			}
		}
	}
	return true
}

type Amphipod struct {
	index    int
	x        int
	y        int
	char     uint8
	moveCost int
}

type State struct {
	pods   []Amphipod
	layout [][]uint8
}

func main() {
	execute.Run(run, tests, puzzle, true)
}

type NodeQ struct {
	Items []Move
	Lock  sync.RWMutex
}

type Move struct {
	Distance int
	state    State
	oldMove  *Move
}

func (s *NodeQ) Enqueue(t Move) {
	s.Lock.Lock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		s.Lock.Unlock()
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.Distance < v.Distance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]Move{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
	//s.items = append(s.items, t)
	s.Lock.Unlock()
}

// Dequeue removes a Node from the start of the queue
func (s *NodeQ) Dequeue() *Move {
	s.Lock.Lock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.Lock.Unlock()
	return &item
}

//NewQ Creates New Queue
func (s *NodeQ) NewQ() *NodeQ {
	s.Lock.Lock()
	s.Items = []Move{}
	s.Lock.Unlock()
	return s
}

// IsEmpty returns true if the queue is empty
func (s *NodeQ) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQ) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}
