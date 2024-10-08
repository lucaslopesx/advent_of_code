package main

import (
	_ "embed"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	in := strings.Split(input, "\n")
	calculateSum(in)
}

func calculateSum(in []string) int {
	res := 0
	for _, card := range in {
		res += calculateScratchcardsSum(card)
	}
	print(res)
	return res
}

func calculateScratchcardsSum(card string) int {
	card = strings.Split(card, ":")[1]
	numbers := strings.Split(card, "|")
	if len(numbers) != 2 {
		return 0
	}

	winNums := getNums(numbers[0])
	choicesNums := getNums(numbers[1])

	sum := 0
	for winNum := range winNums {
		_, exists := choicesNums[winNum]
		if exists {
			if sum == 0 {
				sum += 1
				continue
			}
			sum *= 2
		}
	}

	return sum
}

func getNums(col string) map[int]bool {
	nums := make(map[int]bool)
	var numBuilder strings.Builder
	for i, v := range col {
		if unicode.IsDigit(v) {
			numBuilder.WriteRune(v)
			if i != len(col)-1 {
				continue
			}
		}

		if numBuilder.Len() > 0 {
			num, err := strconv.Atoi(numBuilder.String())
			if err != nil {
				return nil
			}

			nums[num] = true
			numBuilder.Reset()
		}
	}

	return nums
}
