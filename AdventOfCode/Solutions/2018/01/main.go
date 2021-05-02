package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main(){
	pathPrefix, _ := filepath.Abs("AdventOfCode/Solutions/2018/01/input.txt")
	file, _ := os.Open(pathPrefix)
	dat, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println("File reading error", err)
		return
	}

	freq := strings.Split(string(dat), "\r\n")
	computeA(freq)
	computeB(freq)
}

func computeA(freq []string){
	currentFreq := 0

	for _, change := range freq {
		number, _ := strconv.Atoi(change)
		currentFreq += number
	}

	fmt.Printf("answer A: %d\n",  currentFreq)
}

func computeB(freq []string){
	currentFreq := 0

	dict := make(map[int]bool)
	dict[0] = true
	for true{
		for _, change := range freq {
			number, _ := strconv.Atoi(change)
			currentFreq += number
			if dict[currentFreq]{
				fmt.Printf("answer B: %d\n",  currentFreq)
				return
			}
			dict[currentFreq] = true
		}
	}
}