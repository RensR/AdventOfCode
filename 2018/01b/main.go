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
	pathPrefix, _ := filepath.Abs("2018/01b/input.txt")
	file, _ := os.Open(pathPrefix)
	dat, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println("File reading error", err)
		return
	}

	freq := strings.Split(string(dat), "\r\n")


	currentFreq := 0

	dict := make(map[int]bool)
	dict[0] = true
	for _, change := range freq {
		number, _ := strconv.Atoi(change)
		currentFreq += number
		if dict[currentFreq] == true{
			break
		}
		dict[currentFreq] = true
	}

	fmt.Print(currentFreq)
}