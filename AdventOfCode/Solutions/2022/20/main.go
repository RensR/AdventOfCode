package main

import (
	list2 "container/list"
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 20: Grove Positioning System ---
func run(input string) (interface{}, interface{}) {
	var orderedRefsA, orderedRefsB []*list2.Element
	linkedListA, linkedListB := list2.New(), list2.New()

	for _, digit := range strings.Split(input, "\n") {
		parsedDigit := pkg.MustAtoi(digit)
		linkedListA.PushBack(parsedDigit)
		linkedListB.PushBack(parsedDigit * 811589153)
		orderedRefsA = append(orderedRefsA, linkedListA.Back())
		orderedRefsB = append(orderedRefsB, linkedListB.Back())
	}

	Rotate(linkedListA, orderedRefsA)

	for i := 0; i < 10; i++ {
		Rotate(linkedListB, orderedRefsB)
	}

	return GetGroveNumber(linkedListA), GetGroveNumber(linkedListB)
}

func Rotate(linkedList *list2.List, orderedRefs []*list2.Element) {
	for _, numberRef := range orderedRefs {
		// Makes sure we only get positive values, so we only have to deal with `next` calls
		intValue := ((numberRef.Value.(int))%(linkedList.Len()-1) + linkedList.Len() - 1) % (linkedList.Len() - 1)
		//println(fmt.Sprintf("Next to move %14d | moves right %d | %s", numberRef.Value.(int), intValue, PrintList(linkedList)))

		target := numberRef
		for i := 0; i < intValue; i++ {
			target = GetNext(linkedList, target)
		}
		linkedList.MoveAfter(numberRef, target)
	}
}

func PrintList(list *list2.List) string {
	current := list.Front()
	val := ""
	for current != nil {
		val += fmt.Sprintf("%d,", current.Value.(int))
		current = current.Next()
	}
	return val
}

func GetGroveNumber(list *list2.List) (sum int) {
	current := list.Front()
	for current.Value.(int) != 0 {
		current = GetNext(list, current)
	}

	for i := 1; i <= 3000; i++ {
		current = GetNext(list, current)

		if i == 1000 || i == 2000 || i == 3000 {
			sum += current.Value.(int)
		}
	}
	return sum
}

func GetNext(list *list2.List, element *list2.Element) *list2.Element {
	nextRef := element.Next()
	if nextRef == nil {
		nextRef = list.Front()
	}
	return nextRef
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
