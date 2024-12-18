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

func Day2Part2() {
	file, err := os.Open("day2/example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		report := scanner.Text()
		levels := utility.SliceMap(strings.Split(report, " "), func(s string) int64 {
			prevLevel, _ := strconv.ParseInt(s, 10, 64)
			return prevLevel
		})

		diffs := make([]int64, len(levels)-1)
		for idx := 1; idx < len(levels); idx++ {
			diffs[idx-1] = levels[idx] - levels[idx-1]
		}

		fmt.Println()
	}
	fmt.Println(safe)
}

func checkLevels(diffs []int64, skipIdx int, idx int, ascending bool) bool {

}
