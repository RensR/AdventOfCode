package main

import (
	"adventOfCode/helpers"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 22: Reactor Reboot ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	var actions []action

	for _, line := range lines {
		onOff := strings.Split(line, " ")
		xyz := strings.Split(onOff[1], ",")
		x := strings.Split(xyz[0][2:], "..")
		y := strings.Split(xyz[1][2:], "..")
		z := strings.Split(xyz[2][2:], "..")
		a := action{
			on:   onOff[0][1] == 'n',
			from: coord{x: int64(pkg.MustAtoi(x[0])), y: int64(pkg.MustAtoi(y[0])), z: int64(pkg.MustAtoi(z[0]))},
			to:   coord{x: int64(pkg.MustAtoi(x[1])), y: int64(pkg.MustAtoi(y[1])), z: int64(pkg.MustAtoi(z[1]))},
		}
		actions = append(actions, a)
	}

	return initializeReboot(actions), rebootReactor(actions)
}

func countNonOverlapping(actions []action) (volume int64) {
	for _, a := range actions {
		if a.on {
			volume += (a.to.x + 1 - a.from.x) * (a.to.y + 1 - a.from.y) * (a.to.z + 1 - a.from.z)
		}
	}
	return volume
}

func rebootReactor(actions []action) int64 {
	var safeActions []action
	for ; len(actions) > 0; actions = actions[1:] {
		a := actions[0]
		if !a.on {
			continue
		}
		for next := 1; next < len(actions); next++ {
			b := actions[next]
			if a.overlaps(b) {
				removedFirst := actions[1:]
				end := append(a.split(b), removedFirst[next:]...)
				full := append(removedFirst[:next-1], end...)
				return countNonOverlapping(safeActions) + rebootReactor(full)
			}
		}
		safeActions = append(safeActions, a)
	}
	return countNonOverlapping(safeActions)
}

func initializeReboot(actions []action) (on int64) {
	field := make(map[coord]bool)
	for _, a := range actions {
		if a.from.x <= -50 || a.from.x > 50 || a.from.y < -50 || a.from.y > 50 || a.from.z < -50 || a.from.z > 50 {
			continue
		}
		for x := a.from.x; x <= a.to.x; x++ {
			for y := a.from.y; y <= a.to.y; y++ {
				for z := a.from.z; z <= a.to.z; z++ {
					field[coord{
						x: x,
						y: y,
						z: z,
					}] = a.on
				}
			}
		}
	}

	for _, f := range field {
		if f {
			on++
		}
	}
	return on
}

type action struct {
	on   bool
	from coord
	to   coord
}

func (a action) overlaps(b action) bool {
	if a.from.x <= b.from.x && a.to.x >= b.from.x || b.from.x <= a.from.x && b.to.x >= a.from.x {
		if a.from.y <= b.from.y && a.to.y >= b.from.y || b.from.y <= a.from.y && b.to.y >= a.from.y {
			if a.from.z <= b.from.z && a.to.z >= b.from.z || b.from.z <= a.from.z && b.to.z >= a.from.z {
				return true
			}
		}
	}
	return false
}

func (a action) split(b action) (actions []action) {
	overlap := action{
		on: b.on,
		from: coord{
			x: helpers.Max64(a.from.x, b.from.x),
			y: helpers.Max64(a.from.y, b.from.y),
			z: helpers.Max64(a.from.z, b.from.z),
		},
		to: coord{
			x: helpers.Min64(a.to.x, b.to.x),
			y: helpers.Min64(a.to.y, b.to.y),
			z: helpers.Min64(a.to.z, b.to.z),
		},
	}
	actions = append(actions, overlap)
	// x piece before overlap
	if a.from.x < overlap.from.x {
		actions = append(actions, action{
			on:   a.on,
			from: a.from,
			to:   coord{x: overlap.from.x - 1, y: a.to.y, z: a.to.z},
		})
	} else if b.from.x < overlap.from.x {
		actions = append(actions, action{
			on:   b.on,
			from: b.from,
			to:   coord{x: overlap.from.x - 1, y: b.to.y, z: b.to.z},
		})
	}
	// x piece after overlap
	if a.to.x > overlap.to.x {
		actions = append(actions, action{
			on:   a.on,
			from: coord{x: overlap.to.x + 1, y: a.from.y, z: a.from.z},
			to:   a.to,
		})
	} else if b.to.x > overlap.to.x {
		actions = append(actions, action{
			on:   b.on,
			from: coord{x: overlap.to.x + 1, y: b.from.y, z: b.from.z},
			to:   b.to,
		})
	}

	// y piece before overlap
	if a.from.y < overlap.from.y {
		actions = append(actions, action{
			on:   a.on,
			from: coord{x: overlap.from.x, y: a.from.y, z: a.from.z},
			to:   coord{x: overlap.to.x, y: overlap.from.y - 1, z: a.to.z},
		})
	} else if b.from.y < overlap.from.y {
		actions = append(actions, action{
			on:   b.on,
			from: coord{x: overlap.from.x, y: b.from.y, z: b.from.z},
			to:   coord{x: overlap.to.x, y: overlap.from.y - 1, z: b.to.z},
		})
	}
	// y piece after overlap
	if a.to.y > overlap.to.y {
		actions = append(actions, action{
			on:   a.on,
			from: coord{x: overlap.from.x, y: overlap.to.y + 1, z: a.from.z},
			to:   coord{x: overlap.to.x, y: a.to.y, z: a.to.z},
		})
	} else if b.to.y > overlap.to.y {
		actions = append(actions, action{
			on:   b.on,
			from: coord{x: overlap.from.x, y: overlap.to.y + 1, z: b.from.z},
			to:   coord{x: overlap.to.x, y: b.to.y, z: b.to.z},
		})
	}

	// z piece before overlap
	if a.from.z < overlap.from.z {
		actions = append(actions, action{
			on:   a.on,
			from: coord{x: overlap.from.x, y: overlap.from.y, z: a.from.z},
			to:   coord{x: overlap.to.x, y: overlap.to.y, z: overlap.from.z - 1},
		})
	} else if b.from.z < overlap.from.z {
		actions = append(actions, action{
			on:   b.on,
			from: coord{x: overlap.from.x, y: overlap.from.y, z: b.from.z},
			to:   coord{x: overlap.to.x, y: overlap.to.y, z: overlap.from.z - 1},
		})
	}

	// z piece after overlap
	if a.to.z > overlap.to.z {
		actions = append(actions, action{
			on:   a.on,
			from: coord{x: overlap.from.x, y: overlap.from.y, z: overlap.to.z + 1},
			to:   coord{x: overlap.to.x, y: overlap.to.y, z: a.to.z},
		})
	} else if b.to.z > overlap.to.z {
		actions = append(actions, action{
			on:   b.on,
			from: coord{x: overlap.from.x, y: overlap.from.y, z: overlap.to.z + 1},
			to:   coord{x: overlap.to.x, y: overlap.to.y, z: b.to.z},
		})
	}
	return actions
}

type coord struct {
	x int64
	y int64
	z int64
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
