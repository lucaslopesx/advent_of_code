package main

import (
	_ "embed"
	"fmt"
	"sort"
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
	var left []int
	var right []int

	for _, line := range input {
		nums, err := getNums(line)
		if err != nil {
			panic("")
		}
		left = append(left, nums[0])
		right = append(right, nums[1])
	}

	rightMap := getRightListRecurrence(right)

	return sumRightSideSimilarity(left, rightMap)
}

func sumRightSideSimilarity(left []int, rightMap map[int]int) int {
	sort.Ints(left)

	sum := 0
	for _, v := range left {
		recurrence, exists := rightMap[v]
		if !exists {
			continue
		}

		sum += v * recurrence
	}

	return sum
}

func getRightListRecurrence(right []int) map[int]int {
	rightMap := make(map[int]int)
	for _, v := range right {
		rightMap[v]++
	}

	return rightMap
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
