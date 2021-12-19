package main

import (
	"adventOfCode/helpers"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

type Scanner struct {
	loc     *location
	beacons []*location
}

type location struct {
	x int
	y int
	z int
}

// --- Day 19: Beacon scanner ---
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	var scanners []Scanner
	var scanner = Scanner{}
	for _, line := range lines {
		if line == "" {
			scanners = append(scanners, scanner)
			continue
		}
		if line[0] == '-' && line[1] == '-' {
			scanner = Scanner{
				loc:     nil,
				beacons: []*location{},
			}
			continue
		}
		coords := strings.Split(line, ",")
		scanner.beacons = append(scanner.beacons, &location{
			x: pkg.MustAtoi(coords[0]),
			y: pkg.MustAtoi(coords[1]),
			z: pkg.MustAtoi(coords[2]),
		})
	}
	scanners = append(scanners, scanner)

	scanners[0].loc = &location{x: 0, y: 0, z: 0}

	located := []Scanner{scanners[0]}
	scanners = scanners[1:]

	for i := 0; i < len(located); i++ {
		for j := 0; j < len(scanners); j++ {
			if getMatchLevel(&located[i], &scanners[j]) {
				located = append(located, scanners[j])
				scanners = append(scanners[:j], scanners[j+1:]...)
				j--
			}
		}
	}

	return countBeacons(located)
}

func countBeacons(scanners []Scanner) (int, int) {
	allLocations := map[location]bool{}
	maxDist := 0

	for _, scanner := range scanners {
		for _, beacon := range scanner.beacons {
			allLocations[location{
				x: beacon.x + scanner.loc.x,
				y: beacon.y + scanner.loc.y,
				z: beacon.z + scanner.loc.z,
			}] = true
		}
		for _, sc2 := range scanners {
			dist := helpers.Max(scanner.loc.x, sc2.loc.x) - helpers.Min(scanner.loc.x, sc2.loc.x) +
				helpers.Max(scanner.loc.y, sc2.loc.y) - helpers.Min(scanner.loc.y, sc2.loc.y) +
				helpers.Max(scanner.loc.z, sc2.loc.z) - helpers.Min(scanner.loc.z, sc2.loc.z)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	return len(allLocations), maxDist
}

func getMatchLevel(sa *Scanner, sb *Scanner) bool {
	if sa.loc == nil {
		return false
	}

	for xRot := 0; xRot < 4; xRot++ {
		for yRot := 0; yRot < 4; yRot++ {
			for zRot := 0; zRot < 4; zRot++ {
				if isMatch(sa, sb) {
					return true
				}
				// rotate z axis
				rotate(sb, 3)
			}
			// rotate y axis
			rotate(sb, 2)
		}
		// rotate x axis
		rotate(sb, 1)
	}
	return false
}

/*
90 around the Z-axis would be
    |  0   −1   0| |x|   | -y |
    |  1    0   0| |y| = |  x |
    |  0    0   1| |z|   |  z |

90 around the Y-axis would be
    |  0    0   1| |x|   |  z |
    |  0    1   0| |y| = |  y |
    | −1    0   0| |z|   | −x |

90 around the X-axis would be
    |  1    0   0| |x|   |  x |
    |  0    0  −1| |y| = | −z |
    |  0    1   0| |z|   |  y |
*/
func rotate(scanner *Scanner, axis int) {
	for _, beacon := range scanner.beacons {
		if axis == 1 {
			beacon.y, beacon.z = -beacon.z, beacon.y
		} else if axis == 2 {
			beacon.x, beacon.z = beacon.z, -beacon.x
		} else if axis == 3 {
			beacon.x, beacon.y = -beacon.y, beacon.x
		}
	}
}

func isMatch(sa *Scanner, sb *Scanner) bool {
	for _, aBeacon := range sa.beacons {
		for _, bBeacon := range sb.beacons {
			translationVector := getTranslationVector(aBeacon, bBeacon)
			matching := 0
			for _, ab := range sa.beacons {
				for _, bb := range sb.beacons {
					if ab.x == bb.x+translationVector.x && ab.y == bb.y+translationVector.y && ab.z == bb.z+translationVector.z {
						matching++
					}
				}
			}
			if matching >= 2 || translationVector.x == 68 || translationVector.x == -68 {
				matching += 0
			}
			if matching >= 12 {
				newLoc := location{
					x: sa.loc.x + translationVector.x,
					y: sa.loc.y + translationVector.y,
					z: sa.loc.z + translationVector.z,
				}
				sb.loc = &newLoc
				return true
			}
		}
	}
	return false
}

func getTranslationVector(a *location, b *location) location {
	return location{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
