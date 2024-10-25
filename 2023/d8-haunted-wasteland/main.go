package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// type Network struct {
// 	Value string
// 	Left  string
// 	Right string
// }

type Network map[string][]string

func main() {
	in := strings.Split(input, "\n")
	res := runInstructions(in)
	fmt.Print(res)
}

func runInstructions(in []string) int {
	instructions := getInstructions(in[0])
	networks := getNetworks(in[2:])

	return calculateInstructionsSteps(instructions, networks)
}

func calculateInstructionsSteps(instructions []rune, networks Network) int {
	currentValue := "AAA"
	goal := "ZZZ"
	i := 0
	steps := 0
	for {
		instruction := instructions[i]
		switch instruction {
		case 'R':
			currentValue = networks[currentValue][1]
		case 'L':
			currentValue = networks[currentValue][0]
		}

		steps++
		i++

		if i == len(instructions) {
			i = 0
		}

		if currentValue == goal {
			return steps
		}
	}
}

func getInstructions(line string) []rune {
	var instructions []rune
	for _, v := range line {
		if v == '\r' {
			continue
		}
		instructions = append(instructions, v)
	}
	return instructions
}

func getNetworks(lines []string) Network {
	networks := make(map[string][]string)
	for _, line := range lines {
		network, err := extractNetwork(line)
		if err != nil {
			panic("error extracting network")
		}
		for k, v := range network {
			networks[k] = append(networks[k], v...)
		}
	}

	return networks
}

func extractNetwork(line string) (Network, error) {
	lineSplit := strings.Split(line, "=")
	if len(lineSplit) != 2 {
		return Network{}, fmt.Errorf("invalid line size")
	}

	value := strings.TrimSpace(lineSplit[0])
	_ = value

	left := ""
	right := ""
	flag := false

	lineSplit[1] = strings.TrimSpace(lineSplit[1])

	for _, v := range lineSplit[1] {
		if v == '(' {
			continue
		}

		if v == ')' {
			break
		}

		if v == ' ' {
			continue
		}

		if v == ',' {
			flag = true
			continue
		}

		if !flag {
			left += string(v)
			continue
		}

		right += string(v)
	}

	network := make(map[string][]string)
	sides := []string{left, right}
	network[value] = sides

	return network, nil
}
