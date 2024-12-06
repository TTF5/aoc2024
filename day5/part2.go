package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5Part2() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rules = make(map[int][]pageRule)
	var updates [][]int
	var updateMaps []map[int]int

	scanSection := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanSection++
		} else if scanSection == 0 {
			split := strings.Split(line, "|")
			page, _ := strconv.Atoi(split[0])
			rule, _ := strconv.Atoi(split[1])
			if rules[page] == nil {
				rules[page] = make([]pageRule, 0)
			}
			rules[page] = append(rules[page], pageRule{page, rule})
		} else if scanSection == 1 {
			split := strings.Split(line, ",")
			var update []int
			updateMap := make(map[int]int)
			for i, s := range split {
				page, _ := strconv.Atoi(s)
				update = append(update, page)
				updateMap[page] = i + 1
			}
			updateMaps = append(updateMaps, updateMap)
			updates = append(updates, update)
		}
	}

	for k, v := range rules {
		fmt.Printf("%v: %v\n", k, v)
	}

	sum := 0
	for i, update := range updates {
		fmt.Println("== Check rule ===")
		fmt.Printf("%v\n", update)
		if !checkRules(rules, update, updateMaps[i]) {
			fixed := fixOrdering(rules, update, updateMaps[i])
			score := calcUpdateScore(fixed)
			sum += score
		}
	}
	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func moveBefore(pos int, move int, array []int) []int {
	if pos >= move {
		return array
	}
	temp := make([]int, len(array))

	for i := 0; i < pos; i++ {
		temp[i] = array[i]
	}
	temp[pos] = array[move]
	for i := pos + 1; i < move+1; i++ {
		temp[i] = array[i-1]
	}
	for i := move + 1; i < len(array); i++ {
		temp[i] = array[i]
	}
	return temp
}

func fixOrdering(rules map[int][]pageRule, update []int, pagePositions map[int]int) []int {
	for pageIdx, page := range update {
		for _, rule := range rules[page] {
			if rule.before == 0 || rule.page == 0 {
				continue
			}
			if pagePositions[rule.before] == 0 {
				continue
			}
			if pageIdx > pagePositions[rule.before]-1 {
				update = moveBefore(pagePositions[rule.before]-1, pageIdx, update)
				for i, page := range update {
					pagePositions[page] = i + 1
				}
				fmt.Printf("Fixed: %v\n", update)
				return fixOrdering(rules, update, pagePositions)
			}
		}
	}
	return update
}
