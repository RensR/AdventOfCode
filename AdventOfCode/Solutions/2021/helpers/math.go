package helpers

import "strconv"

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FromBitString(bitString string) int64 {
	ox, _ := strconv.ParseInt(bitString, 2, 64)
	return ox
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
