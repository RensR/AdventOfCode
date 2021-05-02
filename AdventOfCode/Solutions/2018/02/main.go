package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main(){
	pathPrefix, _ := filepath.Abs("AdventOfCode/Solutions/2018/02/input.txt")
	file, _ := os.Open(pathPrefix)
	dat, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println("File reading error:", err)
		return
	}

	boxes := strings.Split(string(dat), "\r\n")
	computeA(boxes)
	computeB(boxes)
}

func computeA(boxes []string){
	twice := 0
	trice := 0

	for _, box := range boxes {
		m := make(map[rune]int)
		for _, char := range box {
			m[char] += 1
		}
		for r := range m {
			if m[r] == 3 {
				trice += 1
				break
			}
		}
		for r := range m {
			if m[r] == 2 {
				twice += 1
				break
			}
		}
	}

	fmt.Print(twice * trice)
}

func computeB(boxes []string){
	sort.Strings(boxes)

	for i, box := range boxes {
		nextBox := boxes[i + 1]
		index := -1
		for i, _ := range box {
			if box[i] != nextBox[i] {
				if index != -1 {
					index = -1
					break
				}
				index = i
			}
		}

		if index > -1 {
			fmt.Println("\r\n" + string(delChar([]rune(box), index)))
			break
		}
	}
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}