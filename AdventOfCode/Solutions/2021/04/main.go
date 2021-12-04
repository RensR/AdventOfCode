package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	numbers, boards := parse(input, "\n")

	return playGames(numbers, boards), playToLose(numbers, boards)
}

func playGames(numbers []int, boards []board) int {
	for _, num := range numbers {
		for _, board := range boards {
			won, result := board.checkBoard(num)
			if won {
				return result
			}
		}
	}
	return -1
}

func playToLose(numbers []int, boards []board) int {
	wonBoards := make(map[int]bool)
	for _, num := range numbers {
		for i, board := range boards {
			won, result := board.checkBoard(num)
			if won {
				wonBoards[i] = true
				if len(wonBoards) == len(boards) {
					return result
				}
			}

		}
	}
	return -1
}

type board struct {
	rows    []map[int]bool
	columns map[int]map[int]bool
}

func (b *board) appendRow(line string) {
	numbers := ParseIntList(line, " ")
	row := make(map[int]bool)
	for i, num := range numbers {
		row[num] = true
		if b.columns[i] == nil {
			b.columns[i] = make(map[int]bool)
		}
		b.columns[i][num] = true
	}
	b.rows = append(b.rows, row)
}

func (b *board) checkBoard(number int) (bool, int) {
	hit := false
	for _, row := range b.rows {
		if row[number] {
			hit = true
			delete(row, number)
			if len(row) == 0 {
				return true, b.sumLeft() * number
			}
		}
	}
	if hit {
		for _, column := range b.columns {
			if column[number] {
				delete(column, number)
				if len(column) == 0 {
					return true, b.sumLeft() * number
				}
			}
		}
	}

	return false, 0
}

func (b *board) sumLeft() (sum int) {
	for _, row := range b.rows {
		for number, _ := range row {
			sum += number
		}
	}
	return
}

func newBoard() board {
	return board{
		rows:    []map[int]bool{},
		columns: make(map[int]map[int]bool),
	}
}

func parse(s string, sep string) (numbers []int, boards []board) {
	lines := strings.Split(s, sep)
	board := newBoard()
	for i, line := range lines {
		if i == 0 {
			numbers = pkg.ParseIntList(line, ",")
		} else if line == "" {
			if i > 1 {
				boards = append(boards, board)
				board = newBoard()
			}
		} else {
			board.appendRow(line)
		}
	}
	boards = append(boards, board)
	return
}

func ParseIntList(s, sep string) []int {
	lines := strings.Split(s, sep)
	var list []int
	for _, line := range lines {
		if line != "" {
			nb, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			list = append(list, nb)
		}
	}
	return list
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
