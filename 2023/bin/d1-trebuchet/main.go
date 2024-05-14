package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var stringToNumberMap = map[string]int{
	"nine":  	9,
	"eight": 	8,
	"seven": 	7,
	"six":   	6,
	"five":  	5,
	"four":  	4,
	"three": 	3,
	"two":   	2,
	"one":   	1,
}


//go:embed input.txt
var input string

func main() {
	in := strings.Split(input, "\n")
	res := calculateSum(in)
	fmt.Print(res)
}

func findFirstNumber(message string) (int, error) {
	buffer := ""
	for _, v := range message {
		if unicode.IsDigit(rune(v)) {
			num, err := strconv.Atoi(string(v))
			if err != nil {
				return 0, err
			}
			return num, nil
		}

		buffer += string(v)
		for word, num := range stringToNumberMap {
			if strings.Contains(buffer, word) {
				return num, nil
			}
		}
	}

	return 0, nil
}

func findLastNumber(message string) (int, error) {
	buffer := ""
	for i := len(message) - 1; i >= 0; i--  {
		buffer = string(message[i]) + buffer
		for word, num := range stringToNumberMap {
			if strings.Contains(buffer, word) {
				return num, nil
			}
		}

		if unicode.IsDigit(rune(message[i])) {
			num, err := strconv.Atoi(string(message[i]))
			if err != nil {
				return 0, err
			}
			return num, nil
		}
	}

	return 0, nil
}


func calculateSum(messages []string) int {
	res := 0
	for _, message := range messages {
		firstNumber, err := findFirstNumber(message)
		if err != nil {
			fmt.Println("Error finding the first number", err)
			continue
		}

		lastNumber, err := findLastNumber(message)
		if err != nil {
			fmt.Println("Error finding the last number", err)
			continue
		}

	 	sum := firstNumber * 10 + lastNumber
		res += sum
	}

	return res
}