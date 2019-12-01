package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main(){
	pathPrefix, _ := filepath.Abs("2019/01/input.txt")
	file, _ := os.Open(pathPrefix)
	dat, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println("File reading error", err)
		return
	}

	dataRows := strings.Split(string(dat), "\r\n")

	fuel := 0

	for _, change := range dataRows {
		number, _ := strconv.Atoi(change)
		fuel += int(math.Floor(float64(number/3))) - 2
	}

	fmt.Print(fuel)
}