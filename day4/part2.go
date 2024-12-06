package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var corners = [...]pos{{1, -1}, {1, 1}, {-1, 1}, {-1, -1}}

func Day4Part2() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var startAs []pos

	for row, line := range lines {
		for col, char := range line {
			if char == 'A' {
				startAs = append(startAs, pos{col, row})
			}
		}
	}

	found := 0
	for _, startPos := range startAs {
		if searchXMas(startPos, lines) {
			found++
		}
	}

	fmt.Println(found)
}

// 1  2  3  4
// MM SS MS SM
// SS MM MS SM

func searchXMas(start pos, lines []string) bool {
	// 1
	return getChar(addPos(start, corners[0]), lines) == 'M' &&
		getChar(addPos(start, corners[1]), lines) == 'S' &&
		getChar(addPos(start, corners[2]), lines) == 'S' &&
		getChar(addPos(start, corners[3]), lines) == 'M' ||
		// 2
		getChar(addPos(start, corners[0]), lines) == 'S' &&
			getChar(addPos(start, corners[1]), lines) == 'M' &&
			getChar(addPos(start, corners[2]), lines) == 'M' &&
			getChar(addPos(start, corners[3]), lines) == 'S' ||
		// 3
		getChar(addPos(start, corners[0]), lines) == 'S' &&
			getChar(addPos(start, corners[1]), lines) == 'S' &&
			getChar(addPos(start, corners[2]), lines) == 'M' &&
			getChar(addPos(start, corners[3]), lines) == 'M' ||
		// 3
		getChar(addPos(start, corners[0]), lines) == 'M' &&
			getChar(addPos(start, corners[1]), lines) == 'M' &&
			getChar(addPos(start, corners[2]), lines) == 'S' &&
			getChar(addPos(start, corners[3]), lines) == 'S'
}

func addPos(a, b pos) pos {
	return pos{a.x + b.x, a.y + b.y}
}

func getChar(start pos, lines []string) rune {
	if start.x < 0 || start.x >= len(lines[0]) || start.y < 0 || start.y >= len(lines) {
		return ' '
	}
	return rune(lines[start.y][start.x])
}
