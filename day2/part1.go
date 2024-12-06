package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/TTF5/AoC2024/utility"
)

func Day2Part1() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")

		var prevLevel int64 = 0
		seriesType := 0 // 0 = unset 1 = decreasing 2 = increasing

		isSafe := true

		for idx, level := range levels {

			if idx == 0 {
				prevLevel, _ = strconv.ParseInt(level, 10, 64)
				continue
			}

			levelNum, _ := strconv.ParseInt(level, 10, 64)

			diff := utility.IAbs(prevLevel - levelNum)

			if diff < 1 || diff > 3 {
				isSafe = false
				break
			}

			if prevLevel > levelNum {
				if seriesType == 0 {
					seriesType = 1
				} else if seriesType == 2 {
					isSafe = false
					break
				}

			} else if prevLevel < levelNum {
				if seriesType == 0 {
					seriesType = 2
				} else if seriesType == 1 {
					isSafe = false
					break
				}

			} else {
				isSafe = false
				break
			}
			prevLevel = levelNum
		}
		if isSafe {
			safe++
		}
	}
	fmt.Println(safe)
}
