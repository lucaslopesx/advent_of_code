package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	in := strings.Split(input, "\n")
	res := calculateSum(in)
	fmt.Print(res)
}

func calculateSum(messages []string) int {
	res := 0
	for _, message := range messages {
		firstDigit := 0
		secondDigit := 0
		end := len(message) - 1

		for i := 0; i < len(message); i++ {
			if(firstDigit == 0) {
				v, err := strconv.Atoi(string(message[i]))
				if err != nil {
					continue
				}
				firstDigit = v
			}

			if(secondDigit == 0){
				v, err := strconv.Atoi(string(message[end]))
				if err != nil {
					end--
					continue
				}
				secondDigit = v
			}

			if(firstDigit != 0 && secondDigit != 0){
				break
			}
		}

	 	sum := firstDigit * 10 + secondDigit
		res += sum
	}

	return res
}