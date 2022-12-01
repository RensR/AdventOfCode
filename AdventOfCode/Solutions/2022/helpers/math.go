package helpers

import "strconv"

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Abs[T numbers](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func FromBitString(bitString string) int64 {
	ox, _ := strconv.ParseInt(bitString, 2, 64)
	return ox
}

func Min[T numbers](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T numbers](a, b T) T {
	if a > b {
		return a
	}
	return b
}
