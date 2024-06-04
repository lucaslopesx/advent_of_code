package main

import (
	_ "embed"
	"strconv"
	"strings"
	"unicode"
)

type symbols struct {
	above   map[int]bool
	current map[int]bool
	below   map[int]bool
}

//go:embed input.txt
var input string

func main() {
	in := strings.Split(input, "\n")
	sum := calculateSumInEngineSchematics(in)
	println(sum)
}

func calculateSumInEngineSchematics(in []string) int {
	var symbols_pos = getAllSymbolsPositions(in)

	sum := 0

	var s symbols

	for i, line := range in {
		if i > 0 {
			s.above = symbols_pos[i-1]
		}

		s.below = symbols_pos[i]
		if i != len(in)-1 {
			s.below = symbols_pos[i+1]
		}

		s.current = symbols_pos[i]
		lineSum, err := findValidNumbers(line, s)
		if err != nil {
			print(err)
		}

		sum += lineSum
	}

	return sum
}

func findValidNumbers(line string, s symbols) (int, error) {
	if s.above == nil && s.below == nil {
		return 0, nil
	}

	var sum int
	var num_position []int
	var num_buffer string
	for i, v := range line {

		if unicode.IsDigit(v) {
			_, err := strconv.Atoi(string(v))
			if err != nil {
				return 0, nil
			}

			num_position = append(num_position, i)
			num_buffer += string(v)
			continue
		}

		if len(num_buffer) > 0 {
			num, err := strconv.Atoi(num_buffer)
			if err != nil {
				return 0, err
			}

			if verifyAdjacentSymbol(num_position, s) {
				sum += num
			}
		}

		num_position = nil
		num_buffer = ""
	}

	if len(num_buffer) > 0 {
		num, err := strconv.Atoi(num_buffer)
		if err != nil {
			return 0, err
		}

		if verifyAdjacentSymbol(num_position, s) {
			sum += num
		}
	}

	return sum, nil
}

func verifyAdjacentSymbol(num_position []int, s symbols) bool {
	for i, pos := range num_position {
		if s.above[pos] || s.below[pos] || s.current[pos] {
			return true
		}

		if i == 0 {
			if s.above[pos-1] || s.below[pos-1] || s.current[pos-1] {
				return true
			}
		}

		if i == len(num_position)-1 {
			if s.above[pos+1] || s.below[pos+1] || s.current[pos+1] {
				return true
			}
		}
	}

	return false
}

func getAllSymbolsPositions(in []string) []map[int]bool {
	var res []map[int]bool
	for _, line := range in {
		res = append(res, findLineSymbolsPosition(line))
	}
	return res
}

func findLineSymbolsPosition(line string) map[int]bool {
	m := make(map[int]bool)
	for i, v := range line {
		if unicode.IsSymbol(v) || unicode.IsPunct(v) && string(v) != "." {
			m[i] = true
			continue
		}

		m[i] = false
	}

	return m
}
