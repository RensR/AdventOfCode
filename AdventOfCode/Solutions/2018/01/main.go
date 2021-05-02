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

	currentFreq := 0

	for _, change := range freq {
		number, _ := strconv.Atoi(change)
		currentFreq += number
	}

	fmt.Print(currentFreq)
}