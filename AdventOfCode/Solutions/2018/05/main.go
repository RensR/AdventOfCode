package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
)

func main() {
	pathPrefix, _ := filepath.Abs("AdventOfCode/Solutions/2018/05/input.txt")
	file, _ := os.Open(pathPrefix)
	dat, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Printf("answer A: %d\n", solveA([]rune(string(dat))))

	min := math.MaxInt32
	puzzleInput := []rune(string(dat))
	for char := 65; char < 97; char++ {
		var copyList []rune
		for _, element := range puzzleInput {
			if int(element) != char && int(element) != char+32 {
				copyList = append(copyList, element)
			}
		}
		result := solveA(copyList)
		if result < min {
			min = result
		}
	}
	fmt.Printf("answer B: %d\n", min)
}

func solveA(input []rune) int {
	index := 0
	for {
		if int(input[index])-int(input[index+1]) == 32 ||
			int(input[index])-int(input[index+1]) == -32 {
			input = delTwoChars(input, index)
			if index != 0 {
				index -= 1
			}
		} else {
			index += 1
		}

		if index >= len(input)-1 {
			return len(input)
		}
	}
}

func delTwoChars(s []rune, index int) []rune {
	return append(s[0:index], s[index+2:]...)
}
