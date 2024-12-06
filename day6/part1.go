package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/TTF5/AoC2024/utility"
)

const (
	TILE_OBSTACLE = '#'
	TILE_EMPTY    = '.'
	TILE_VISITED  = 'X'
)

type Guard struct {
	position  utility.Point
	direction int
}

func parseMap(input string) ([][]rune, Guard) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var roomMap [][]rune

	guard := Guard{position: utility.Point{X: 0, Y: 0}, direction: 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for x, tile := range line {
			if tile == '^' {
				guard = Guard{position: utility.Point{X: x, Y: len(roomMap)}, direction: 0}
				tile = TILE_VISITED
			}
		}
		roomMap = append(roomMap, []rune(line))
	}
	return roomMap, guard
}

func printMap(roomMap [][]rune) {
	println("====================================")
	for _, row := range roomMap {
		for _, tile := range row {
			print(string(tile))
		}
		println()
	}
}

func doStep(roomMap [][]rune, guard Guard) ([][]rune, Guard) {
	pos := guard.position
	nextPos := pos

	switch guard.direction {
	case 0: // Up
		nextPos = utility.Point{X: pos.X, Y: pos.Y - 1}
	case 1: // Right
		nextPos = utility.Point{X: pos.X + 1, Y: pos.Y}
	case 2: // Down
		nextPos = utility.Point{X: pos.X, Y: pos.Y + 1}
	case 3: // Left
		nextPos = utility.Point{X: pos.X - 1, Y: pos.Y}
	}

	if posOnMap(roomMap, nextPos) == false {
		guard.position = nextPos
		return roomMap, guard
	}

	if roomMap[nextPos.Y][nextPos.X] == TILE_OBSTACLE {
		guard.direction = (guard.direction + 1) % 4
	} else {
		roomMap[nextPos.Y][nextPos.X] = TILE_VISITED
		guard.position = nextPos
	}

	return roomMap, guard
}

func posOnMap(roomMap [][]rune, pos utility.Point) bool {
	return pos.X >= 0 && pos.X < len(roomMap[0]) && pos.Y >= 0 && pos.Y < len(roomMap)
}

func countVisited(roomMap [][]rune) int {
	count := 0
	for _, row := range roomMap {
		for _, tile := range row {
			if tile == TILE_VISITED {
				count++
			}
		}
	}
	return count
}

func Day6Part1() {
	roomMap, guard := parseMap("day6/input.txt")

	fmt.Println(guard)
	// printMap(roomMap)

	for posOnMap(roomMap, guard.position) {
		roomMap, guard = doStep(roomMap, guard)
		// printMap(roomMap)
	}

	visited := countVisited(roomMap)

	fmt.Println("Visited: ", visited)
}
