package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	DIR_UP = iota
	DIR_UPRIGHT
	DIR_RIGHT
	DIR_DOWNRIGHT
	DIR_DOWN
	DIR_DOWNLEFT
	DIR_LEFT
	DIR_UPLEFT
)

var dirs = [...]int{DIR_UP, DIR_UPRIGHT, DIR_RIGHT, DIR_DOWNRIGHT, DIR_DOWN, DIR_DOWNLEFT, DIR_LEFT, DIR_UPLEFT}

const searchStr = "XMAS"

type pos struct {
	x int
	y int
}

func Day4Part1() {
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

	var startXs []pos

	for row, line := range lines {
		for col, char := range line {
			if char == 'X' {
				startXs = append(startXs, pos{col, row})
			}
		}
	}

	found := 0
	for _, startPos := range startXs {
		for dir := range dirs {
			if search(startPos, dir, lines) {
				found++
			}
		}
	}

	fmt.Println(found)
}

func getDirVect(dir int) pos {
	switch dir {
	case DIR_UP:
		return pos{0, -1}
	case DIR_UPRIGHT:
		return pos{1, -1}
	case DIR_RIGHT:
		return pos{1, 0}
	case DIR_DOWNRIGHT:
		return pos{1, 1}
	case DIR_DOWN:
		return pos{0, 1}
	case DIR_DOWNLEFT:
		return pos{-1, 1}
	case DIR_LEFT:
		return pos{-1, 0}
	case DIR_UPLEFT:
		return pos{-1, -1}
	}
	log.Fatal("Invalid direction")
	return pos{0, 0}
}

func search(start pos, dir int, lines []string) bool {
	vect := getDirVect(dir)
	for _, searchChar := range searchStr {
		if start.x < 0 || start.x >= len(lines[0]) || start.y < 0 || start.y >= len(lines) {
			return false
		}
		var currChar rune = rune(lines[start.y][start.x])
		if currChar != searchChar {
			return false
		}
		start.x += vect.x
		start.y += vect.y
	}
	return true
}
