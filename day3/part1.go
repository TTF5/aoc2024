package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day3Part1() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fmt.Println(match)
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			fmt.Println(a, b)
			sum += a * b
		}
	}
	fmt.Println(sum)
}
