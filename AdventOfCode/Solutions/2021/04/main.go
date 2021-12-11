package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 4: Giant Squid ---
func run(input string) (interface{}, interface{}) {
	numbers, bingoBoards := parse(strings.Split(input, "\n"))
	return playGames(numbers, bingoBoards), playToLose(numbers, bingoBoards)
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
	return sum
}

func newBoard() board {
	return board{
		rows:    []map[int]bool{},
		columns: make(map[int]map[int]bool),
	}
}

func parse(lines []string) (numbers []int, boards []board) {
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
	return numbers, append(boards, board)
}

func ParseIntList(s, sep string) (result []int) {
	lines := strings.Split(s, sep)
	for _, line := range lines {
		if line != "" {
			result = append(result, pkg.MustAtoi(line))
		}
	}
	return result
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
