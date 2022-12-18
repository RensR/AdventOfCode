package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

const leftWall, rightWall = uint32(1077952576), uint32(0x01010101)

type Snapshot struct {
	Shape     int
	WindIndex int
	Tower     uint64
}

type SnapResult struct {
	BlockNumber int
	Height      int
}

// --- Day 17: Pyroclastic Flow ---
func run(input string) (interface{}, interface{}) {
	tower := NewTower()

	cache := make(map[Snapshot]SnapResult)
	heightAdd := 0
	a1 := 0
	for {
		tower = DropBlock(tower, input)
		if tower.BlockNumber > 1e4 && heightAdd == 0 {
			snapshot := Snapshot{
				Shape:     tower.BlockNumber % 5,
				WindIndex: tower.WindIndex,
				Tower:     tower.Get8AtHeight(tower.BlockNumber),
			}
			if old, ok := cache[snapshot]; ok {
				blocksInCycle := tower.BlockNumber - old.BlockNumber
				heightIncrease := tower.Height() - old.Height
				cyclesNeeded := (1e12 - tower.BlockNumber) / blocksInCycle

				tower.BlockNumber += cyclesNeeded * blocksInCycle
				heightAdd = heightIncrease * cyclesNeeded
			} else {
				cache[snapshot] = SnapResult{
					BlockNumber: tower.BlockNumber,
					Height:      tower.Height(),
				}
			}
		}

		if tower.BlockNumber == 1e12 {
			return a1, tower.Height() + heightAdd
		}

		if tower.BlockNumber == 2022 {
			a1 = tower.Height()
		}
	}
}

type Tower struct {
	Blocks      []uint8
	WindIndex   int
	BlockNumber int
}

func NewTower() Tower {
	return Tower{
		Blocks:      []uint8{},
		WindIndex:   0,
		BlockNumber: 0,
	}
}

func (t *Tower) Get8AtHeight(height int) uint64 {
	return uint64(t.GetAtHeight(height))<<32 + uint64(t.GetAtHeight(height-4))
}

func (t *Tower) GetAtHeight(height int) uint32 {
	if height > len(t.Blocks) {
		return 0
	}

	result := uint32(0)
	for i := 0; i < 4; i++ {
		if height+i >= len(t.Blocks) {
			continue
		}
		if height+i < 0 {
			result += 255 << (8 * i)
			continue
		}
		currentBlock := uint32(t.Blocks[height+i])
		result += currentBlock << (8 * i)
	}
	return result
}

func (t *Tower) Height() int {
	return len(t.Blocks)
}

func (t *Tower) Draw() {
	for i := t.Height() - 1; i >= 0; i-- {
		println(DrawSingleLine(t.Blocks[i]))
	}
}

func DrawFourLines(lines uint32) string {
	allLines := ""
	for i := 3; i >= 0; i-- {
		shapeLayer := uint8(lines >> (8 * (3 - i)))
		line := DrawSingleLine(shapeLayer)
		allLines = line + "\n" + allLines
	}
	return allLines
}

func DrawSingleLine(line uint8) string {
	drawLine := "|"
	for j := 1; j < 8; j++ {
		if line&1 == 0 {
			drawLine = "." + drawLine
		} else {
			drawLine = "@" + drawLine
		}
		line = line >> 1
	}
	return "|" + drawLine
}

func DropBlock(tower Tower, wind string) Tower {
	currentHeight := tower.Height() + 3
	shape := getShape(tower.BlockNumber)
	tower.BlockNumber++

	for {
		// Handle wind
		currentWind := wind[tower.WindIndex]
		tower.WindIndex++
		if tower.WindIndex >= len(wind) {
			tower.WindIndex = 0
		}

		shape = HandleWind(currentWind, shape, tower.GetAtHeight(currentHeight))

		//towerSilhouette := tower.GetAtHeight(currentHeight)
		//println(fmt.Sprintf("Tower at height %d", currentHeight))
		//towerSilhouetteDown := tower.GetAtHeight(currentHeight - 4)
		//println(DrawFourLines(towerSilhouette|shape) + DrawFourLines(towerSilhouetteDown))

		// Handle falling
		if currentHeight > tower.Height()+1 {
			currentHeight--
			continue
		}
		towerAtNextIteration := tower.GetAtHeight(currentHeight - 1)
		wouldCollide := shape&towerAtNextIteration != 0
		if currentHeight == 0 || wouldCollide {
			// Update the tower with the new block
			towerHeightBeforeInsert := tower.Height()
			for offset := 0; offset < 4; offset++ {
				if currentHeight+offset >= towerHeightBeforeInsert {
					shapeLayer := uint8(shape >> (8 * offset))
					if shapeLayer != 0 {
						tower.Blocks = append(tower.Blocks, shapeLayer)
					}
				} else {
					tower.Blocks[currentHeight+offset] |= uint8(shape >> (8 * offset))
				}
			}

			break
		} else {
			currentHeight--
		}
	}

	return tower
}

func HandleWind(direction uint8, shape uint32, tower uint32) uint32 {
	var movedPiece uint32
	if direction == '>' {
		if shape&rightWall == 0 {
			movedPiece = shape >> 1
		} else {
			return shape
		}
	} else {
		if shape&leftWall == 0 {
			movedPiece = shape << 1
		} else {
			return shape
		}
	}
	if movedPiece&tower == 0 {
		return movedPiece
	}
	return shape
}

func getShape(blockNumber int) uint32 {
	switch blockNumber % 5 {
	case 0:
		// . . . . . . . .
		// . . . . . . . .
		// . . . . . . . .
		// . . . @ @ @ @ .
		return 30
	case 1:
		// . . . . . . . .
		// . . . . @ . . .
		// . . . @ @ @ . .
		// . . . . @ . . .
		return 531464
		// 000010000001110000001000
	case 2:
		// . . . . . . . .
		// . . . . . @ . .
		// . . . . . @ . .
		// . . . @ @ @ . .
		return 263196
		// 000001000000010000011100
	case 3:
		// . . . @ . . . .
		// . . . @ . . . .
		// . . . @ . . . .
		// . . . @ . . . .
		return 269488144
		// 00010000000100000001000000010000
	case 4:
		// . . . . . . . .
		// . . . . . . . .
		// . . . @ @ . . .
		// . . . @ @ . . .
		return 6168
		// 0001100000011000
	default:
		panic("")
	}

}

func main() {
	execute.Run(run, tests, puzzle, true)
}
