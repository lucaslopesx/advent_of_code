package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Network map[string][]string

func main() {
	in := strings.Split(input, "\n")
	res := runInstructions(in)
	fmt.Print(res)
}

func runInstructions(in []string) int {
	instructions := getInstructions(in[0])
	networks, startingPoints := getNetworks(in[2:])

	var steps []int

	for _, start := range startingPoints {
		steps = append(steps, calculateInstructionsSteps(instructions, networks, start))
	}

	result := LCM(steps[0], steps[1], steps[2:]...)

	return result
}

func calculateInstructionsSteps(instructions []rune, networks Network, start string) int {
	currentValue := start
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

		if _, found := strings.CutSuffix(currentValue, "Z"); found {
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

func getNetworks(lines []string) (Network, []string) {
	networks := make(map[string][]string)
	var startingPoints []string
	for _, line := range lines {
		network, startingPoint, err := extractNetwork(line)
		if err != nil {
			panic("error extracting network")
		}

		if startingPoint != "" {
			startingPoints = append(startingPoints, startingPoint)
		}

		for k, v := range network {
			networks[k] = append(networks[k], v...)
		}
	}

	return networks, startingPoints
}

func extractNetwork(line string) (Network, string, error) {
	lineSplit := strings.Split(line, "=")
	startingPoint := ""
	if len(lineSplit) != 2 {
		return Network{}, "", fmt.Errorf("invalid line size")
	}

	value := strings.TrimSpace(lineSplit[0])
	_, exists := strings.CutSuffix(value, "A")
	if exists {
		startingPoint = value
	}
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

	return network, startingPoint, nil
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
