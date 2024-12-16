package day14

import (
	"fmt"
	"os"
	"os/exec"
	"time"
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

func Day14Part2() {
	var robots []Robot
	var width int
	var height int
	seconds := 500
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
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		for i, robot := range robots {
			robots[i] = moveRobot(robot, width, height, 1)
		}
		PrintMap(robots, width, height)
		time.Sleep(100 * time.Millisecond)
	}

}
