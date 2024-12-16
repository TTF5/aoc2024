package day7

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

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

func worker(results chan<- int, eq Equation) {
	fmt.Println("Processing", eq)
	ch := utility.GenerateCombinations([]Operation{OP_ADD, OP_CONCAT, OP_MULT}, len(eq.values)-1)
	for ops := range ch {
		if checkEquation2(eq, ops) {
			results <- eq.result
			break
		}
	}
}

func Day7Part2() {
	eqns := parseInput("day7/input.txt")

	results := make(chan int, len(eqns))
	var wg sync.WaitGroup
	noEqns := len(eqns)

	start := time.Now()

	wg.Add(noEqns)
	for _, eq := range eqns {
		go func(eq Equation) {
			worker(results, eq)
			wg.Done()
		}(eq)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	calibResult := 0
	for val := range results {
		calibResult += val
	}

	fmt.Println(calibResult)

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}
