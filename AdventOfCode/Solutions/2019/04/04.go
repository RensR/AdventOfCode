package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := 0

	for i := 278384; i <= 824795; i++ {
		if checkDigit(i) {
			c++
		}
	}

	fmt.Print(c)
}

func checkDigit(digit int) bool {
	same := false
	input := strconv.Itoa(digit)
	for i, c := range input {
		if i == 0 {
			continue
		}
		if uint8(c) < input[i-1] {
			return false
		} else if uint8(c) == input[i-1] {
			if i+1 >= len(input) || uint8(c) != input[i+1] {
				if i == 1 || uint8(c) != input[i-2] {
					same = true
				}
			}
		}
	}
	return same
}
