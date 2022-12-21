package main

import (
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"golang.org/x/exp/slices"

	math2 "adventOfCode/helpers/math"
)

// --- Day 15: Beacon Exclusion Zone ---
func run(input string) (interface{}, interface{}) {
	var beacons, sensors []image.Point

	lines := strings.Split(input, "\n")
	testIndex := pkg.MustAtoi(lines[0])
	maxSearchSpace := testIndex * 2

	for _, line := range lines[1:] {
		var sx, sy, bx, by int
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil {
			panic(err)
		}
		sensor, beacon := image.Point{X: sx, Y: sy}, image.Point{X: bx, Y: by}
		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}

	ranges := echooooo(beacons, sensors)

	b1 := 0
	for y := 0; y <= maxSearchSpace; y++ {
		if x := FindHoleInRange(ranges[y], maxSearchSpace); x != -1 {
			b1 = 4_000_000*x + y
		}
	}

	return SumRanges(ranges[testIndex]) - 1, b1
}

func FindHoleInRange(line []Range, max int) int {
	slices.SortFunc(line, func(a, b Range) bool {
		return a.start < b.start
	})

	count := line[0].end - line[0].start + 1
	latest := line[0].end
	for i := 1; i < len(line) && latest+1 <= max; i++ {
		current := line[i]
		if current.end <= latest {
			continue
		}
		if current.start > latest+1 {
			return latest + 1
		}

		count += current.end - latest + 1
		latest = current.end
	}

	if latest+1 < max {
		return latest + 1
	}

	return -1
}

func SumRanges(line []Range) int {
	slices.SortFunc(line, func(a, b Range) bool {
		return a.start < b.start
	})

	count := line[0].end - line[0].start + 1
	latest := line[0].end
	for i := 1; i < len(line); i++ {
		current := line[i]
		if current.end < latest {
			continue
		}
		if current.start > latest {
			latest = current.end
			count += current.end - current.start + 1
			continue
		}

		count += current.end - latest
		latest = current.end
	}

	return count
}

type Range struct {
	start, end int
}

func echooooo(beacons, sensors []image.Point) map[int][]Range {
	ranges := make(map[int][]Range)

	for _, sensor := range sensors {
		// get distance to the closest beacon
		distance := math.MaxInt
		for _, beacon := range beacons {
			currentDistance := math2.Abs(sensor.X-beacon.X) + math2.Abs(sensor.Y-beacon.Y)
			if currentDistance < distance {
				distance = currentDistance
			}
		}

		for dy := 0; dy <= distance; dy++ {
			dx := distance - dy

			xRange := Range{start: sensor.X - dx, end: sensor.X + dx}
			ranges[sensor.Y+dy] = append(ranges[sensor.Y+dy], xRange)
			ranges[sensor.Y-dy] = append(ranges[sensor.Y-dy], xRange)
		}
	}

	return ranges
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
