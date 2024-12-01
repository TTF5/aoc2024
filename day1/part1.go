package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func IAbs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1 []int64
	var list2 []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, num1)

		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0

	for idx, num := range list1 {
		dist := IAbs(num - list2[idx])
		sum += int(dist)
	}

	fmt.Println(sum)
}
