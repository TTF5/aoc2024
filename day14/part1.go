package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func parseInput(input string) []Robot {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var robots []Robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse x, y, vx, vy from the line
		parts := strings.Split(line, " ")
		pos := strings.TrimPrefix(parts[0], "p=")
		vel := strings.TrimPrefix(parts[1], "v=")

		posParts := strings.Split(pos, ",")
		velParts := strings.Split(vel, ",")

		x, _ := strconv.Atoi(posParts[0])
		y, _ := strconv.Atoi(posParts[1])
		vx, _ := strconv.Atoi(velParts[0])
		vy, _ := strconv.Atoi(velParts[1])

		robot := Robot{x: x, y: y, vx: vx, vy: vy}

		// Append the Robot struct to the robots slice
		robots = append(robots, robot)
	}
	return robots
}

func moveRobot(robot Robot, width int, height int, seconds int) Robot {
	for i := 0; i < seconds; i++ {
		robot.x += robot.vx
		robot.y += robot.vy
		if robot.x < 0 {
			robot.x += width
		} else if robot.x >= width {
			robot.x -= width
		}
		if robot.y < 0 {
			robot.y += height
		} else if robot.y >= height {
			robot.y -= height
		}
	}
	return robot
}

func test() {
	robot := Robot{x: 2, y: 4, vx: 2, vy: -3}
	width := 11
	height := 7
	fmt.Println(robot)
	fmt.Println(moveRobot(robot, width, height, 1))
	fmt.Println(moveRobot(robot, width, height, 2))
	fmt.Println(moveRobot(robot, width, height, 3))
	fmt.Println(moveRobot(robot, width, height, 4))
	fmt.Println(moveRobot(robot, width, height, 5))
	fmt.Println(moveRobot(robot, width, height, 6))

}

func Day14Part1() {
	var robots []Robot
	var width int
	var height int
	seconds := 100
	example := false

	// test()
	// return

	if example {
		robots = parseInput("day14/example.txt")
		width = 11
		height = 7
	} else {
		robots = parseInput("day14/input.txt")
		width = 101
		height = 103
	}

	fmt.Println(robots, width, height)

	for i, robot := range robots {
		robots[i] = moveRobot(robot, width, height, seconds)
	}

	fmt.Println(robots)

	var quadrants [4]int
	for _, robot := range robots {
		if robot.x < width/2 && robot.y < height/2 {
			quadrants[0]++
		} else if robot.x > width/2 && robot.y < height/2 {
			quadrants[1]++
		} else if robot.x < width/2 && robot.y > height/2 {
			quadrants[2]++
		} else if robot.x > width/2 && robot.y > height/2 {
			quadrants[3]++
		}
	}

}
