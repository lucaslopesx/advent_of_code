package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	in := strings.Split(input, "\n")
	res := calc(in)
	fmt.Println(res)
}

func calc(input []string) int {

	var reports [][]int

	for _, line := range input {
		nums, err := getNums(line)
		if err != nil {
			panic("failed to parse numbers from input")
		}
		reports = append(reports, nums)
	}

	sum := 0
	for _, level := range reports {
		if isLevelSafe(level) {
			sum++
		}
	}

	return sum
}

func isLevelSafe(level []int) bool {
	isIncreasing := level[0] < level[1]
	for i := range level {
		if i == len(level)-1 {
			continue
		}

		diff := level[i] - level[i+1]
		if diff == 0 {
			return false
		}

		if (diff > 0) != !isIncreasing {
			return false
		}

		if !isDiffSafe(diff) {
			return false
		}
	}

	return true
}

func isDiffSafe(diff int) bool {
	allowed := []int{1, 2, 3}
	diff = int(math.Abs(float64(diff)))

	return slices.Contains(allowed, diff)
}

func getNums(input string) ([]int, error) {
	numStrings := strings.Fields(input)
	numbers := make([]int, 0, len(numStrings))
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
