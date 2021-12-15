package main

import (
	"adventOfCode/helpers"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 15: Chiton ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	riskMap := make(map[int]map[int]int)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if riskMap[x] == nil {
				riskMap[x] = make(map[int]int)
			}
			riskMap[x][y] = pkg.MustAtoi(string(lines[y][x]))
		}
	}

	cost := bfs(riskMap, len(lines[0]), len(lines))
	bigMap := expandMap(riskMap, len(lines[0])-1, len(lines)-1)
	return cost, bfs(bigMap, len(lines[0])*5-1, len(lines)*5-1)
}

func expandMap(riskMap map[int]map[int]int, maxX int, maxY int) map[int]map[int]int {
	newMap := make(map[int]map[int]int)

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for i, risks := range riskMap {
				xx := maxX*x + i
				if _, ok := newMap[xx]; !ok {
					newMap[xx] = make(map[int]int)
				}
				for ii, risk := range risks {
					if risk+x+y > 9 {
						newMap[xx][maxY*y+ii] = risk + x + y - 9
					} else {
						newMap[xx][maxY*y+ii] = risk + x + y
					}
				}
			}
		}
	}
	return newMap
}

func bfs(riskMap map[int]map[int]int, goalX int, goalY int) int {
	Q := helpers.NodeQueue{}
	Q.NewQ()
	Q.Enqueue(helpers.Vertex{X: 1, Y: 0, Distance: riskMap[1][0]})
	Q.Enqueue(helpers.Vertex{X: 0, Y: 1, Distance: riskMap[0][1]})
	visited := map[helpers.Vertex]bool{{X: 0, Y: 0}: true}

	for !Q.IsEmpty() {
		active := Q.Dequeue()
		if active.X == goalX && active.Y == goalY {
			return active.Distance
		}
		if visited[helpers.Vertex{X: active.X, Y: active.Y}] {
			continue
		}
		visited[helpers.Vertex{X: active.X, Y: active.Y}] = true
		for _, direction := range helpers.UpDownLeftRight {
			yy, xx := active.Y+direction.Y, active.X+direction.X
			// Out of bounds check
			if yy < 0 || yy >= goalY || xx < 0 || xx >= goalX || visited[helpers.Vertex{X: xx, Y: yy}] {
				continue
			}
			Q.Enqueue(helpers.Vertex{X: xx, Y: yy, Distance: riskMap[xx][yy] + active.Distance})
		}
	}

	return -1
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
