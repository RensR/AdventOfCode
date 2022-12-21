package _map

import "golang.org/x/exp/constraints"

func MapMaxValue[T constraints.Ordered, K comparable](m map[K]T) T {
	max := T(0)
	for _, value := range m {
		if value > max {
			max = value
		}
	}
	return max
}
