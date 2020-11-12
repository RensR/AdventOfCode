package main

import "fmt"

func main(){
	input := []uint {
		1, 1, 0, 0, 1,
		0, 0, 0, 1, 1,
		0, 1, 0, 1, 1,
		1, 0, 0, 1, 0,
		0, 0, 1, 0, 0}

	found := make(map[uint]bool)

	for true {
		diversity := calculate(input)

		if found[diversity] {
			fmt.Print(diversity)
			return
		}

		found[diversity] = true
		input = mutate(input)
	}
}


func mutate(planet []uint)(newPlanet []uint){
	newPlanet = append([]uint(nil), planet...)

	for area := uint32(0); area < 25; area++ {
		influence := uint(0)
		// left
		if area % 5 != 0 {
			influence += planet[area - 1]
		}
		// right
		if area % 5 != 4 {
			influence += planet[area + 1]
		}
		// up
		if area > 4 {
			influence += planet[area - 5]
		}
		// down
		if area < 20 {
			influence += planet[area + 5]
		}

		if planet[area] == 1 && influence != 1 {
			newPlanet[area] = 0
		} else if planet[area] == 0 && (influence == 1 || influence == 2){
			newPlanet[area] = 1
		}
	}

	return
}

func calculate(planet []uint)(diversity uint){
	diversity = 0

	for area := uint32(0); area < 25; area++ {
		diversity += planet[area] << area
	}

	return
}

