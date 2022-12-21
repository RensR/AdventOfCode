package datastructures

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

func (stack *Stack[T]) Len() int {
	return len(stack.items)
}

func (stack *Stack[T]) Pop() (value T) {
	n := len(stack.items)
	if n <= 0 {
		panic("Cannot pop an empty stack!")
	}
	value = stack.items[n-1]
	stack.items = stack.items[:n-1]
	return
}
