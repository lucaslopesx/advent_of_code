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
	in := strings.Split(input, "\n")
	res := calculateWaysOfBeatingRaces(in)
	fmt.Println(res)
}

func calculateWaysOfBeatingRaces(in []string) int {
	if len(in) != 2 {
		return 0
	}

	times := getNums(in[0])
	distances := getNums(in[1])
	if len(times) != len(distances) {
		return 0
	}

	ways := make(map[int]int)
	for i, time := range times {
		goal := distances[i]
		remaining := time
		for hold := 0; hold < time; hold++ {
			if (hold * remaining) > goal {
				w := hold * 2
				ways[i] = time - w + 1
				break
			}

			remaining--
		}
	}

	total := 1
	for _, v := range ways {
		total *= v
	}

	return total
}

func getNums(line string) []int {
	var nums []int
	var numBuilder strings.Builder
	for i, v := range line {
		if unicode.IsDigit(v) {
			numBuilder.WriteRune(v)
		}

		if i != len(line)-1 {
			continue
		}

		if numBuilder.Len() > 0 {
			num, err := strconv.Atoi(numBuilder.String())
			if err != nil {
				return nil
			}

			nums = append(nums, num)
			numBuilder.Reset()
		}
	}

	return nums
}
