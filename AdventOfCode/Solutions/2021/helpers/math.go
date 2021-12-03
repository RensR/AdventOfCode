package helpers

import "strconv"

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func FromBitString(bitString string) int64 {
	ox, _ := strconv.ParseInt(bitString, 2, 64)
	return ox
}
