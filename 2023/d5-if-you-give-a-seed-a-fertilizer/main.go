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

type Position struct {
	destination int
	source      int
	length      int
}

func main() {
	in := strings.Split(input, "\n")

	res := calculateClosestLocation(in)

	print(res)
}

func calculateClosestLocation(in []string) int {
	seeds := getNums(in[0])
	maps := getMaps(in[2:])

	var locations []int
	for _, seed := range seeds {
		loc := seed
		for i := 0; i < len(maps); i++ {
			for _, position := range maps[i] {
				x, err := position.tryGetNextDestination(loc)
				if err == nil {
					loc = x
					break
				}
			}
		}
		locations = append(locations, loc)
	}

	closestLocation := locations[0]
	for _, loc := range locations {
		if loc < closestLocation {
			closestLocation = loc
		}
	}

	return closestLocation
}

func getNums(line string) []int {
	var nums []int
	var numBuilder strings.Builder
	for i, v := range line {
		if unicode.IsDigit(v) {
			numBuilder.WriteRune(v)
			if i != len(line)-1 {
				continue
			}
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

func getPosition(line string) (Position, error) {
	nums := getNums(line)

	if len(nums) != 3 {
		return Position{}, fmt.Errorf("Position not found")
	}

	return Position{
		destination: nums[0],
		source:      nums[1],
		length:      nums[2],
	}, nil
}

func getMaps(in []string) map[int][]Position {
	currentMap := 0
	var maps = make(map[int][]Position)
	var positions []Position
	for _, v := range in {
		if v == "\r" || v == "" {
			if len(positions) > 0 {
				maps[currentMap] = positions
				positions = []Position{}
			}
			currentMap++
		}

		position, err := getPosition(v)
		if err == nil {
			positions = append(positions, position)
		}
	}

	if len(positions) > 0 {
		maps[currentMap] = positions
	}

	return maps
}

func (position Position) tryGetNextDestination(x int) (int, error) {
	if x < position.source {
		return x, fmt.Errorf("x value lower then source position")
	}

	if x >= (position.source + position.length) {
		return x, fmt.Errorf("position not found")
	}

	diff := position.destination - position.source

	return (x + diff), nil
}
