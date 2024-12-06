package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pageRule struct {
	page   int
	before int
}

func Day5Part1() {
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
		if checkRules(rules, update, updateMaps[i]) {
			score := calcUpdateScore(update)
			sum += score
			fmt.Println("Valid - Score: " + strconv.Itoa(score))
		} else {
			fmt.Println("Invalid")
		}
	}
	fmt.Println("Sum: " + strconv.Itoa(sum))
}

func calcUpdateScore(update []int) int {
	return update[len(update)/2]
}

func checkRules(rules map[int][]pageRule, update []int, pagePositions map[int]int) bool {
	for pageIdx, page := range update {
		for _, rule := range rules[page] {
			if rule.before == 0 || rule.page == 0 { // Rule not found!
				fmt.Println("Rule not found for page: " + strconv.Itoa(page))
				continue
			}
			if pagePositions[rule.before] == 0 {
				continue
			}
			if pageIdx > pagePositions[rule.before]-1 {
				fmt.Println("Rule failed: " + strconv.Itoa(rule.before) + " was before " + strconv.Itoa(rule.page))
				return false
			}
		}
	}
	return true
}
