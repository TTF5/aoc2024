package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/TTF5/AoC2024/utility"
)

type Equation struct {
	values []int
	result int
}

const (
	OP_ADD    Operation = iota
	OP_MULT   Operation = iota
	OP_CONCAT Operation = iota
)

type Operation int8

func parseInput(input string) []Equation {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var eqns []Equation

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		valuesStr := strings.Trim(split[1], " ")

		result, _ := strconv.Atoi(split[0])
		values := utility.SliceMap(strings.Split(valuesStr, " "), func(s string) int {
			val, _ := strconv.Atoi(s)
			return val
		})

		eq := Equation{values: values, result: result}

		eqns = append(eqns, eq)
	}
	return eqns
}

func checkEquation(eq Equation, ops []Operation) bool {
	result := eq.values[0]

	for i, val := range eq.values[1:] {
		if ops[i] == OP_ADD {
			result += val
		} else if ops[i] == OP_MULT {
			result *= val
		}
	}

	return result == eq.result
}

func Day7Part1() {
	eqns := parseInput("day7/input.txt")

	calibResult := 0

	for _, eq := range eqns {
		ch := utility.GenerateCombinations([]Operation{OP_ADD, OP_MULT}, len(eq.values)-1)
		for ops := range ch {
			if checkEquation(eq, ops) {
				calibResult += eq.result
				break
			}
		}
	}

	fmt.Println(calibResult)
}
