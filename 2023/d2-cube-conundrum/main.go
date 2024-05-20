package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var availableCubes = map[string]int{
	"red":  		12,
	"green": 		13,
	"blue": 		14,
}

//go:embed input.txt
var input string

func findPlayableGameId(game string) (int, error) {
	availableCubeBuffer := ""

	currentCubeValueBuffer := ""
	currentCubeValue := 0

	gameIdBuffer := ""
	gameId := 0


	for _, v := range game {

		if(gameId == 0) {
			if unicode.IsDigit(v) {
				_, err := strconv.Atoi(string(v))
				if err != nil {
					return 0, err
				}

				gameIdBuffer += string(v)
				continue
			}

			if len(gameIdBuffer) > 0 {
				num, err := strconv.Atoi(gameIdBuffer)
				if err != nil {
					return 0, err
				}

				gameId = num
			}
			
			continue
		}


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
		for cube, availableValue := range availableCubes {
			if strings.Contains(availableCubeBuffer, cube) {
				if currentCubeValue > availableValue {
					return 0, errors.New("error current cube value is higher than the available value")
				}
				currentCubeValue = 0
				currentCubeValueBuffer = ""
				availableCubeBuffer = ""
				continue
			}
		}
	}

	return gameId, nil
}

func calculateSumOfGameIds(games []string) int {
	sumOfGameIds := 0
	
	for _, game := range games {
		gameId, err := findPlayableGameId(game)
		if err != nil {
			fmt.Println("error finding playable game id", err)
			continue
		}
		sumOfGameIds += gameId
	}

	return sumOfGameIds
}

func main() {
	games := strings.Split(input, "\n")
	sum := calculateSumOfGameIds(games)
	fmt.Print(sum)
}