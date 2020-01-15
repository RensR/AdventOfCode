package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	perm := permutations([]int{0,1,2,3,4})

	max := 0
	for _, permutation := range perm{
		output := 0
		for _, val := range permutation{
			output = intCode(val, output)
		}

		if max < output{
			max = output
		}

	}
 	fmt.Print(strconv.Itoa(max))
}

func intCode(inputVar int, prevOutput int)(output int){
	//input := []int {3,23,3,24,1002,24,10,24,1002,23,-1,23, 101,5,23,23,1,24,23,23,4,23,99,0,0}
	input := []int {3,8,1001,8,10,8,105,1,0,0,21,34,59,76,101,114,195,276,357,438,99999,3,9,1001,9,4,9,1002,9,4,9,4,9,99,3,9,102,4,9,9,101,2,9,9,102,4,9,9,1001,9,3,9,102,2,9,9,4,9,99,3,9,101,4,9,9,102,5,9,9,101,5,9,9,4,9,99,3,9,102,2,9,9,1001,9,4,9,102,4,9,9,1001,9,4,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,99}
	maxArgLen := 5
	for i := 0; i <= len(input); {
		optCode := input[i] % 100
		modes := LeftPad2Len(strconv.Itoa(input[i]), "0", maxArgLen)
		arg1, arg2, arg3 := 0, 0, 0
		if optCode == 1 || optCode == 2 || optCode == 4 || optCode == 5|| optCode == 6|| optCode == 7|| optCode == 8 {
			if modes[maxArgLen - 3] == '1'{
				arg1 = input[i + 1]
			} else{
				arg1 = input[input[i + 1]]
			}
		}
		if optCode == 1 || optCode == 2 || optCode == 5|| optCode == 6|| optCode == 7|| optCode == 8 {
			if modes[maxArgLen - 4] == '1'{
				arg2 = input[i + 2]
			} else {
				arg2 = input[input[i + 2]]
			}
		}
		if optCode == 7|| optCode == 8 {
			arg3 = input[i + 3]
		}
		switch optCode {
		case 1:
			input[input[i + 3]] = arg1 + arg2
			i += 4
		case 2:
			input[input[i + 3]] = arg1 * arg2
			i += 4
		case 3:
			input[input[i + 1]] = inputVar
			inputVar = prevOutput
			i += 2
		case 4:
			output = arg1
			i += 2
		case 5:
			if arg1 != 0{
				i = arg2
			} else {
				i += 3
			}
		case 6:
			if arg1 == 0 {
				i = arg2
			} else {
				i += 3
			}
		case 7:
			if arg1 < arg2 {
				input[arg3] = 1
			} else {
				input[arg3] = 0
			}
			i += 4
		case 8:
			if arg1 == arg2 {
				input[arg3] = 1
			} else {
				input[arg3] = 0
			}
			i += 4
		case 99:
			return
		}
	}

	return
}

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func permutations(arr []int)[][]int{
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int){
		if n == 1{
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++{
				helper(arr, n - 1)
				if n % 2 == 1{
					tmp := arr[i]
					arr[i] = arr[n - 1]
					arr[n - 1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n - 1]
					arr[n - 1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}