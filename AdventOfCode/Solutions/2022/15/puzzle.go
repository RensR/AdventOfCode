package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzleTest,
		`26`,
		`56000011`,
	},
	{
		puzzle,
		`4985193`,
		`11583882601918`,
	},
}

var puzzleTest = `10
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

var puzzle = `2000000
Sensor at x=3523437, y=2746095: closest beacon is at x=3546605, y=2721324
Sensor at x=282831, y=991087: closest beacon is at x=743030, y=-87472
Sensor at x=1473740, y=3283213: closest beacon is at x=1846785, y=3045894
Sensor at x=1290563, y=46916: closest beacon is at x=743030, y=-87472
Sensor at x=3999451, y=15688: closest beacon is at x=3283637, y=-753607
Sensor at x=1139483, y=2716286: closest beacon is at x=1846785, y=3045894
Sensor at x=3137614, y=2929987: closest beacon is at x=3392051, y=3245262
Sensor at x=2667083, y=2286333: closest beacon is at x=2126582, y=2282363
Sensor at x=3699264, y=2920959: closest beacon is at x=3546605, y=2721324
Sensor at x=3280991, y=2338486: closest beacon is at x=3546605, y=2721324
Sensor at x=833202, y=92320: closest beacon is at x=743030, y=-87472
Sensor at x=3961416, y=2485266: closest beacon is at x=3546605, y=2721324
Sensor at x=3002132, y=3500345: closest beacon is at x=3392051, y=3245262
Sensor at x=2482128, y=2934657: closest beacon is at x=1846785, y=3045894
Sensor at x=111006, y=2376713: closest beacon is at x=354526, y=3163958
Sensor at x=424237, y=2718408: closest beacon is at x=354526, y=3163958
Sensor at x=3954504, y=3606495: closest beacon is at x=3392051, y=3245262
Sensor at x=2275050, y=2067292: closest beacon is at x=2333853, y=2000000
Sensor at x=1944813, y=2557878: closest beacon is at x=2126582, y=2282363
Sensor at x=2227536, y=2152792: closest beacon is at x=2126582, y=2282363
Sensor at x=3633714, y=1229193: closest beacon is at x=3546605, y=2721324
Sensor at x=1446898, y=1674290: closest beacon is at x=2333853, y=2000000
Sensor at x=3713985, y=2744503: closest beacon is at x=3546605, y=2721324
Sensor at x=2281504, y=3945638: closest beacon is at x=1846785, y=3045894
Sensor at x=822012, y=3898848: closest beacon is at x=354526, y=3163958
Sensor at x=89817, y=3512049: closest beacon is at x=354526, y=3163958
Sensor at x=2594265, y=638715: closest beacon is at x=2333853, y=2000000`