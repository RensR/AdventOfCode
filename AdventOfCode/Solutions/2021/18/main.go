package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 18: Snailfish ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	var totalSumSnail = parseSnailSum(lines[0])
	for i := 1; i < len(lines); i++ {
		totalSumSnail = &snailSum{
			left:  totalSumSnail,
			right: parseSnailSum(lines[i]),
		}
		totalSumSnail.right.parent = totalSumSnail
		totalSumSnail.left.parent = totalSumSnail

		setDepth(totalSumSnail, 0)
		reduce(totalSumSnail)
	}

	largestNum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			sum2 := &snailSum{
				left:  parseSnailSum(lines[j]),
				right: parseSnailSum(lines[i]),
			}
			sum2.right.parent = sum2
			sum2.left.parent = sum2

			setDepth(sum2, 0)
			result := getMagnitude(reduce(sum2))
			if result > largestNum {
				largestNum = result
			}
		}
	}

	result := getMagnitude(totalSumSnail)
	return result, largestNum
}

func setDepth(sum *snailSum, depth int64) {
	sum.depth = depth
	if sum.left != nil {
		setDepth(sum.left, depth+1)
	}
	if sum.right != nil {
		setDepth(sum.right, depth+1)
	}
}

type snailSum struct {
	left   *snailSum
	right  *snailSum
	depth  int64
	lVal   int
	rVal   int
	parent *snailSum
}

func (sum *snailSum) tryExplode() bool {
	if sum.depth >= 4 && sum.left == nil && sum.right == nil {
		addLeft(sum, sum.lVal)
		addRight(sum, sum.rVal)

		if sum.parent.left == sum {
			sum.parent.left = nil
			sum.parent.lVal = 0
		} else {
			sum.parent.right = nil
			sum.parent.rVal = 0
		}
		return true
	} else {
		if sum.left != nil {
			if sum.left.tryExplode() {
				return true
			}
		}
		if sum.right != nil {
			return sum.right.tryExplode()
		}
		return false
	}
}

func addLeft(sum *snailSum, val int) {
	if sum.parent == nil {
		return
	}

	if sum.parent.left == sum {
		addLeft(sum.parent, val)
	} else {
		if sum.parent.left == nil {
			sum.parent.lVal += val
		} else {
			getRightLeaf(sum.parent.left).rVal += val
		}
	}
}

func addRight(sum *snailSum, val int) {
	if sum.parent == nil {
		return
	}
	if sum.parent.left == sum {
		if sum.parent.right == nil {
			sum.parent.rVal += val
		} else {
			getLeftLeaf(sum.parent.right).lVal += val
		}
	} else {
		addRight(sum.parent, val)
	}
}

func getLeftLeaf(sum *snailSum) *snailSum {
	current := sum
	for current.left != nil {
		current = current.left
	}
	return current
}

func getRightLeaf(sum *snailSum) *snailSum {
	current := sum
	for current.right != nil {
		current = current.right
	}
	return current
}

func (sum *snailSum) trySplit() bool {
	if sum == nil {
		return false
	}
	if sum.lVal > 9 {
		sum.left = &snailSum{
			depth:  sum.depth + 1,
			lVal:   sum.lVal / 2,
			rVal:   (sum.lVal + 1) / 2,
			parent: sum,
		}
		sum.lVal = 0
		return true
	} else if sum.left.trySplit() {
		return true
	} else if sum.rVal > 9 {
		sum.right = &snailSum{
			depth:  sum.depth + 1,
			lVal:   sum.rVal / 2,
			rVal:   (sum.rVal + 1) / 2,
			parent: sum,
		}
		sum.rVal = 0
		return true
	}
	return sum.right.trySplit()
}

func reduce(sum *snailSum) *snailSum {
	for {
		if !sum.tryExplode() {
			if !sum.trySplit() {
				return sum
			}
		}
	}
}

func parseSnailSum(math string) *snailSum {
	if len(math) == 5 {
		leftRightVal := strings.Split(math[1:len(math)-1], ",")
		return &snailSum{
			lVal: pkg.MustAtoi(leftRightVal[0]),
			rVal: pkg.MustAtoi(leftRightVal[1]),
		}
	}

	numOpenBrackets := 0
	for i := 0; i < len(math); i++ {
		if math[i] == ',' && numOpenBrackets == 1 {
			var result = &snailSum{}
			// single digit
			if i-1 == 1 {
				result.lVal = pkg.MustAtoi(math[1:i])
			} else { // snail
				result.left = parseSnailSum(math[1:i])
				result.left.parent = result
			}
			if len(math)-1-(i+1) == 1 {
				result.rVal = pkg.MustAtoi(math[i+1 : len(math)-1])
			} else {
				result.right = parseSnailSum(math[i+1 : len(math)-1])
				result.right.parent = result
			}

			return result
		}
		if math[i] == '[' {
			numOpenBrackets++
		} else if math[i] == ']' {
			numOpenBrackets--
		}
	}
	panic(1)
}

func getMagnitude(sum *snailSum) int {
	left := sum.lVal
	if sum.left != nil {
		left = getMagnitude(sum.left)
	}
	if sum.right == nil {
		return left*3 + sum.rVal*2
	}
	return left*3 + getMagnitude(sum.right)*2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
