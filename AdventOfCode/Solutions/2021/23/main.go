package main

import (
	"adventOfCode/helpers"
	"fmt"
	"strings"
	"sync"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var letters = []uint8{'A', 'B', 'C', 'D'}

// --- Day 23: Amphipod ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	moveCost := map[int32]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

	var layout = make([][]uint8, len(lines[0]))
	var pods [8]Amphipod
	index := 0
	for y, line := range lines {
		for x, c := range line {
			layout[x] = append(layout[x], uint8(c))
			switch c {
			case 'A', 'B', 'C', 'D':
				pods[index] = Amphipod{
					index:    index,
					x:        x,
					y:        y,
					char:     uint8(c),
					moveCost: moveCost[c],
				}
				index++
			}
		}
	}

	Q := NodeQ{}
	Q.NewQ()
	currentState := &Move{Distance: 0, state: State{pods, layout, 4}}
	seenStates := make(map[string]bool)
	for !currentState.state.IsFinal() {
		moves := currentState.state.getValidMoves(seenStates, currentState.Distance)
		for _, move := range moves {
			move.oldMove = currentState
			Q.Enqueue(move)
		}

		currentState = Q.Dequeue()
	}
	printing := currentState
	for printing != nil {
		(*printing).state.Print()
		printing = (*printing).oldMove
	}
	return currentState.Distance, run2(lines)
}

func run2(lines []string) int64 {

	return 0
}

func (state State) Print() string {
	totalString := ""
	for y := 0; y < 4; y++ {
		line := ""
		for x, _ := range state.layout {
			line += string(state.layout[x][y])
		}
		fmt.Println(line)
		line += totalString
	}
	return totalString
}

func (state State) GetSignature() string {
	totalString := ""
	for y := 0; y < 4; y++ {
		line := ""
		for x, _ := range state.layout {
			line += string(state.layout[x][y])
		}
		totalString += line
	}
	return totalString
}

func (state State) getValidMoves(lookup map[string]bool, currentDistance int) (moves []Move) {
	for _, pod := range state.pods {
		if pod.IsDone(state) {
			continue
		}

		seenThisPod := map[helpers.Location]bool{helpers.Location{X: pod.x, Y: pod.y}: true}
		newMoves := DoStep(state, pod, 0, seenThisPod, pod.y < 2)
		for _, move := range newMoves {
			signature := move.state.GetSignature()
			if lookup[signature] && !move.state.IsFinal() {
				continue
			}
			lookup[signature] = true
			//move.state.Print()
			move.Distance += currentDistance
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
				if yy == 1 || pod.y == 3 { // do another step, never stop here
					moves = append(moves, DoStep(newState, newState.pods[pod.index], steps, seenThisPod, fromHallway)...)
				} else { // final home checks, if the room isn't theirs don't use it
					if xx != getXForPodId(pod.char) {
						continue
					}
					if state.layout[xx][3] == pod.char {
						// Done with the entire letter, move to y = 2
						moves = append(moves, Move{
							Distance: steps * pod.moveCost,
							state:    newState,
						})
					} else if state.layout[xx][3] == '.' {
						// Done with this pod, move to y = 3
						steps++
						newState.layout[xx][2] = '.'
						newState.layout[xx][3] = pod.char
						moves = append(moves, Move{
							Distance: steps * pod.moveCost,
							state:    newState,
						})
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
	newState := State{pods: [8]Amphipod{}, layout: makeCopyLayout(state.layout)}
	newState.layout[pod.x][pod.y] = '.'
	newState.layout[xx][yy] = pod.char
	for _, amphi := range state.pods {
		copyAmphi := amphi
		if copyAmphi.index == pod.index {
			copyAmphi.x = xx
			copyAmphi.y = yy
		}
		newState.pods[amphi.index] = copyAmphi
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

func (pod Amphipod) IsDone(state State) bool {
	x := getXForPodId(pod.char)
	return pod.x == x && (pod.y == 3 || pod.y == 2 && state.layout[x][3] == pod.char)
}

func getXForPodId(podId uint8) int {
	return int((podId-'A')*2) + 3
}

func (state State) IsFinal() bool {
	for _, letter := range letters {
		x := getXForPodId(letter)

		for i := 2; i < state.height+2; i++ {

		}
		if state.layout[x][2] != letter || state.layout[x][3] != letter {
			return false
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
	pods   [8]Amphipod
	layout [][]uint8
	height int
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
