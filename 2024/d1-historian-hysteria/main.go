package main

import (
	_ "embed"
	"fmt"
	"math"
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

	return sumSideDistances(left, right)
}

func sumSideDistances(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		diffAbs := math.Abs(float64(diff))
		sum += int(diffAbs)
	}

	return sum

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
