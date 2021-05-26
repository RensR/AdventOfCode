package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println([]byte("aA"))
	solveA([]rune(string(dat)))
}

func solveA(input []rune) {
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

		if index >= len(input) - 1 {
			fmt.Printf("answer A: %d\n", len(input))
			return
		}
	}
}
func delTwoChars(s []rune, index int) []rune {
	return append(s[0:index], s[index+2:]...)
}
