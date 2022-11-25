package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"

	"adventOfCode/helpers"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// --- Day 16: ??? ---
func run(input string) (interface{}, interface{}) {
	decoded, _ := hex.DecodeString(input)
	bits := ""
	for _, n := range decoded {
		bits += fmt.Sprintf("%08b", n)
	}
	result, operatorResult, _ := parseBits(bits, 0)
	return result, operatorResult
}

func parseBits(bits string, start int64) (versionSum int64, operatorResult int64, length int64) {
	i := start
	totalVersions := readNextXBits(bits, &i, 3)
	typeID := readNextXBits(bits, &i, 3)
	switch typeID {
	case 4:
		totalValues := ""
		for {
			firstBit := readNextXBits(bits, &i, 1)
			totalValues += bits[i : i+4]
			i += 4
			if firstBit == 0 {
				break
			}
		}
		operatorResult = parseBitsToInt(totalValues)
	default:
		lengthId := readNextXBits(bits, &i, 1)
		var values []int64
		switch lengthId {
		case 0: // total length in bits of subpackets is next 15
			subPacketLength := readNextXBits(bits, &i, 15)
			for currentLength := int64(0); currentLength < subPacketLength; {
				value, opResult, length := parseBits(bits, i)
				i += length
				totalVersions += value
				currentLength += length
				values = append(values, opResult)
			}
		case 1:
			subPacketCount := readNextXBits(bits, &i, 11)
			for subPackets := int64(0); subPackets < subPacketCount; subPackets++ {
				value, opResult, length := parseBits(bits, i)
				i += length
				totalVersions += value
				values = append(values, opResult)
			}
		}
		switch typeID {
		case 0: // sum
			for _, val := range values {
				operatorResult += val
			}
		case 1: // product
			operatorResult = int64(1)
			for _, val := range values {
				operatorResult *= val
			}
		case 2: // min
			operatorResult = math.MaxInt
			for _, val := range values {
				operatorResult = helpers.Min(operatorResult, val)
			}
		case 3: // max
			for _, val := range values {
				operatorResult = helpers.Max(operatorResult, val)
			}
		case 5: // greater
			if values[0] > values[1] {
				operatorResult = int64(1)
			}
		case 6: // less
			if values[0] < values[1] {
				operatorResult = int64(1)
			}
		case 7: // ==
			if values[0] == values[1] {
				operatorResult = int64(1)
			}
		}
	}

	return totalVersions, operatorResult, i - start
}

func readNextXBits(bits string, i *int64, length int64) int64 {
	val := parseBitsToInt(bits[*i : *i+length])
	*i += length
	return val
}

func parseBitsToInt(bits string) int64 {
	if i, err := strconv.ParseInt(bits, 2, 64); err != nil {
		panic(err)
	} else {
		return i
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
