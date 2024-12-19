package day14

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"

	"github.com/sergeymakinen/go-bmp"
)

func PrintMap(robots []Robot, width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			found := false
			for _, robot := range robots {
				if robot.x == x && robot.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func robotsToBitmap(seconds int, robots []Robot, width int, height int) {
	var bitmap = image.NewPaletted(image.Rect(0, 0, width, height), []color.Color{color.Black, color.White})

	for _, robot := range robots {
		bitmap.Set(robot.x, robot.y, color.White)
	}

	filename := "day14/images/second_" + strconv.Itoa(seconds) + ".bmp"

	w, _ := os.Create(filename)
	bmp.Encode(w, bitmap)
}

func Day14Part2() {
	var robots []Robot
	var width int
	var height int
	seconds := 10000
	example := false

	if example {
		robots = parseInput("day14/example.txt")
		width = 11
		height = 7
	} else {
		robots = parseInput("day14/input.txt")
		width = 101
		height = 103
	}

	for s := 0; s < seconds; s++ {
		for i, robot := range robots {
			robots[i] = moveRobot(robot, width, height, 1)
		}
		robotsToBitmap(s, robots, width, height)
	}
}
