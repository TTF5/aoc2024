package day6

import (
	"fmt"

	"github.com/TTF5/AoC2024/utility"
)

func checkRepeatedPosition(route []Guard, guard Guard) bool {
	for _, g := range route {
		if g.position == guard.position && g.direction == guard.direction {
			// fmt.Println(route, guard)
			return true
		}
	}
	return false
}

func checkLoop(inMap [][]rune, guard Guard, obstacle utility.Point) bool {
	route := []Guard{guard}
	loop := false

	roomMap := make([][]rune, len(inMap))
	for i := range inMap {
		roomMap[i] = make([]rune, len(inMap[i]))
		copy(roomMap[i], inMap[i])
	}

	roomMap[obstacle.Y][obstacle.X] = TILE_OBSTACLE

	for posOnMap(roomMap, guard.position) {
		roomMap, guard = doStep(roomMap, guard)
		// printMap(roomMap)
		if checkRepeatedPosition(route, guard) {
			loop = true
			// printMap(roomMap)
			break
		}
		route = append(route, guard)
	}

	return loop
}

func Day6Part2() {
	roomMap, guard := parseMap("day6/input.txt")

	// fmt.Println("Check loop fun:", checkLoop(roomMap, guard, utility.Point{X: 0, Y: 0}))

	loopsFound := 0

	for y, _ := range roomMap {
		for x, _ := range roomMap[y] {
			obstacle := utility.Point{X: x, Y: y}
			if checkLoop(roomMap, guard, obstacle) {
				// fmt.Println("Loop found at", obstacle)
				loopsFound++
			} else {
				// fmt.Println("No Loop", obstacle)
			}
		}
	}

	fmt.Println(loopsFound)
}
