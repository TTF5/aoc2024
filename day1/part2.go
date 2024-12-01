package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var keys []int
	var list []int
	m := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		keys = append(keys, int(num1))

		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, int(num2))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, num := range list {
		m[num] = m[num] + 1
	}

	sum := 0

	for _, num := range keys {
		if m[num] > 0 {
			sum += num * m[num]
		}
	}

	fmt.Println(sum)
}
