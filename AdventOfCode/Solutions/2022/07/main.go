package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

// --- Day 7: No Space Left On Device ---
func run(input string) (interface{}, interface{}) {
	totalAboveTenK := 0

	root := Folder{
		content: map[string]*Folder{},
		parent:  nil,
		size:    0,
	}

	var current *Folder

	for _, line := range strings.Split(input, "\n") {
		commands := strings.Split(line, " ")
		switch commands[0] {
		case "$":
			switch commands[1] {
			case "cd":
				switch commands[2] {
				case "/":
					current = &root
				case "..":
					current = current.parent
				default:
					current = current.content[commands[2]]
				}
			case "ls":
			}
		case "dir":
			current.content[commands[1]] = &Folder{
				content: map[string]*Folder{},
				parent:  current,
				size:    0,
			}
		default:
			tmp := current
			for ; current != nil; current = current.parent {
				current.size += pkg.MustAtoi(commands[0])
			}
			current = tmp
		}
	}

	var todo helpers.Stack[*Folder]

	todo.Push(&root)
	needed := 30000000 - (70000000 - root.size)
	minDiff := 70000000

	for todo.Len() != 0 {
		folder := todo.Pop()
		if folder.size <= 100000 {
			totalAboveTenK += folder.size
		}
		if folder.size >= needed {
			if folder.size < minDiff {
				minDiff = folder.size
			}
		}
		for _, f := range folder.content {
			todo.Push(f)
		}
	}

	return totalAboveTenK, minDiff
}

type Folder struct {
	content map[string]*Folder
	parent  *Folder
	size    int
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
