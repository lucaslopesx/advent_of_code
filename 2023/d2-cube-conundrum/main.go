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

func findPlayableGame(game string) (int, error) {
	var cubes = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	currentCubeValue := 0
	currentCubeValueBuffer := ""
	availableCubeBuffer := ""

	parts := strings.Split(game, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("error invalid game")
	}
	game = parts[1]

	for _, v := range game {
		if currentCubeValue == 0 {
			if unicode.IsDigit(v) {
				_, err := strconv.Atoi(string(v))
				if err != nil {
					return 0, err
				}
				currentCubeValueBuffer += string(v)
				continue
			}

			if len(currentCubeValueBuffer) > 0 {
				num, err := strconv.Atoi(currentCubeValueBuffer)
				if err != nil {
					return 0, err
				}
				currentCubeValue = num
			}
			continue
		}

		availableCubeBuffer += string(v)
		for cube, value := range cubes {
			if strings.Contains(availableCubeBuffer, cube) {
				if currentCubeValue > value {
					cubes[cube] = currentCubeValue
				}
				currentCubeValue = 0
				currentCubeValueBuffer = ""
				availableCubeBuffer = ""
				continue
			}
		}
	}

	res := 1
	for _, v := range cubes {
		res *= v
	}

	return res, nil
}

func calculateSumOfGames(games []string) int {
	sum := 0

	for _, game := range games {
		game, err := findPlayableGame(game)
		if err != nil {
			fmt.Println("error finding playable game", err)
			continue
		}
		sum += game
	}

	return sum
}

func main() {
	games := strings.Split(input, "\n")
	sum := calculateSumOfGames(games)
	fmt.Print(sum)
}
