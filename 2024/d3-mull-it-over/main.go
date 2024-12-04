package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	res := calc(input)
	fmt.Println(res)
}

func calc(input string) int {
	sum := 0
	enabled := true
	for i := range input {
		if i+7 < len(input) && input[i:i+7] == "don't()" {
			enabled = false
			continue
		}

		if i+4 >= len(input) {
			continue
		}

		if input[i:i+4] == "do()" {
			enabled = true
			continue
		}

		if input[i:i+4] == "mul(" && enabled {
			sum += getInstructionMultiplied(input[i+4:])
			continue
		}
	}

	return sum
}

func getInstructionMultiplied(input string) int {
	buffer := ""
	for _, v := range input {
		if v == ' ' {
			return 0
		}

		if v != ',' && v != ')' && !unicode.IsDigit(v) {
			return 0
		}

		if v == ')' {
			nums, err := getNums(buffer)
			if err != nil || len(nums) != 2 {
				return 0
			}
			return nums[0] * nums[1]
		}

		buffer += string(v)
	}
	return 0
}

func getNums(input string) ([]int, error) {
	numStrings := strings.Split(input, ",")
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
