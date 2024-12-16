package day7

import (
	"fmt"
	"strconv"

	"github.com/TTF5/AoC2024/utility"
)

func checkEquation2(eq Equation, ops []Operation) bool {
	result := eq.values[0]

	for i, val := range eq.values[1:] {
		if ops[i] == OP_ADD {
			result += val
		} else if ops[i] == OP_MULT {
			result *= val
		} else if ops[i] == OP_CONCAT {
			resultStr := strconv.Itoa(result) + strconv.Itoa(val)
			result, _ = strconv.Atoi(resultStr)
		}
	}

	return result == eq.result
}

func Day7Part2() {
	eqns := parseInput("day7/input.txt")

	calibResult := 0

	noEqns := len(eqns)

	for i, eq := range eqns {
		ch := utility.GenerateCombinations([]Operation{OP_ADD, OP_CONCAT, OP_MULT}, len(eq.values)-1)
		fmt.Println("Checking equation", i+1, "of", noEqns)
		for ops := range ch {
			if checkEquation2(eq, ops) {
				calibResult += eq.result
				break
			}
		}
	}

	fmt.Println(calibResult)
}
