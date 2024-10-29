package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type History struct {
	Sequences [][]int
}

func main() {
	in := strings.Split(input, "\n")
	res := runReports(in)
	fmt.Print(res)
}

func runReports(in []string) int {
	var histories []History
	for _, line := range in {
		history := processLine(line)
		histories = append(histories, history)
	}

	sum := 0
	for _, v := range histories {
		sum += extrapolateHistory(v.Sequences)
	}
	return sum
}

func extrapolateHistory(sequences [][]int) int {
	sequencesSize := len(sequences) - 1
	lastValue := 0
	for i := sequencesSize; i > 0; i-- {
		current := sequences[i]
		up := sequences[i-1]

		currentValue := 0
		if i != sequencesSize {
			currentValue = current[0]
		}

		upperValue := up[0]

		extrapolatedValue := upperValue - currentValue
		sequences[i-1] = append([]int{extrapolatedValue}, sequences[i-1]...)
		lastValue = extrapolatedValue
	}

	return lastValue
}

func processLine(line string) History {
	var start, err = parseNumbers(line)
	if err != nil {
		return History{}
	}
	var diffs []int
	var sequences [][]int
	sequences = append(sequences, start)
	currentValues := start
	for {
		isAllZeroes := true
		for i := 0; i < len(currentValues)-1; i++ {
			num1 := currentValues[i]
			num2 := 0
			if i != len(currentValues) {
				num2 = currentValues[i+1]
			}

			diff := num2 - num1
			if diff != 0 {
				isAllZeroes = false
			}

			diffs = append(diffs, diff)
		}

		sequences = append(sequences, diffs)
		if isAllZeroes {
			break
		}

		currentValues = diffs
		diffs = nil
	}

	return History{
		Sequences: sequences,
	}
}

func parseNumbers(input string) ([]int, error) {
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
