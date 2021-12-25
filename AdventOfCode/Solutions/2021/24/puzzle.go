package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	//{
	//	puzzle,
	//	``,
	//	``,
	//},
}

var puzzle = `inp w
mul x 0
add x z
mod x 26
div z 1
add x 14
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 1
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 15
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 7
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 15
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 13
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -6
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 14
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 0
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -4
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 13
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 15
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 11
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 15
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 6
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 1
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x 0
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 7
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x 0
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 11
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -3
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 14
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -9
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 4
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -9
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y`

var puzzleBackup = `inp w
add z w
add z 1 
inp w 		// z [2..10]
mul z 26
add z w
add z 7
inp w		// z [x*26 + w + 7]
mul z 26
add z w
add z 13
inp w
div z 26
add x -6
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y
inp w
add x 14
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 0
mul y x
add z y
inp w
div z 26
add x -4
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 13
mul y x
add z y
inp w
add x 15
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 11
mul y x
add z y
inp w
add x 15
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 6
mul y x
add z y
inp w
add x 11
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 1
mul y x
add z y
inp w
div z 26
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 7
mul y x
add z y
inp w
div z 26
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 11
mul y x
add z y
inp w
div z 26
add x -3
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 14
mul y x
add z y
inp w
div z 26
add x -9
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 4
mul y x
add z y
inp w
div z 26
add x -9
eql x w
eql x 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y`
