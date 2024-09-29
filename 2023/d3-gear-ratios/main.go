package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Gear struct {
	row, col int
	parts    []int
}

//go:embed input.txt
var input string

func main() {
	in := strings.Split(input, "\n")
	sum := calculateGearRatioSum(in)
	fmt.Println("Result:", sum)
}

func calculateGearRatioSum(in []string) int {
	schematic := readEngineSchematic(in)
	gears := findGears(schematic)

	sum := 0
	for _, v := range gears {
		ratio := v.parts[0] * v.parts[1]
		sum += ratio
	}
	return sum
}

func readEngineSchematic(in []string) [][]rune {
	schematic := make([][]rune, len(in))
	for i, line := range in {
		schematic[i] = []rune(line)
	}
	return schematic
}

func findGears(schematic [][]rune) []Gear {
	var gears []Gear
	for row := range schematic {
		for col := range schematic[row] {
			if schematic[row][col] == '*' {
				parts := findParts(schematic, row, col)
				if len(parts) != 2 {
					continue
				}

				gear := Gear{
					col:   col,
					row:   row,
					parts: parts,
				}
				gears = append(gears, gear)
			}
		}
	}
	return gears
}

// ...  .->[row-1][col-1] .->[row-1][col] .->[row-1][col+1]
// .*.  .->[row]  [col-1] *->[row]  [col] .->[row]  [col+1]
// ...  .->[row+1][col-1] .->[row+1][col] .->[row+1][col+1]
func findParts(schematic [][]rune, row, col int) []int {
	var parts []int
	var seenPositions = make(map[string]bool)
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}

			newRow, newCol := row+dr, col+dc
			if newRow >= 0 && newRow < len(schematic) && newCol >= 0 && newCol < len(schematic[newRow]) {
				num, position, found := getFullNumber(schematic, newRow, newCol)
				pRow := fmt.Sprintf("%d-%s", newRow, position)
				if found && !seenPositions[pRow] {
					parts = append(parts, num)
					seenPositions[pRow] = true
				}
			}
		}
	}

	return parts
}

func getFullNumber(schematic [][]rune, row, col int) (int, string, bool) {
	line := schematic[row]
	if !unicode.IsDigit(line[col]) {
		return 0, "", false
	}

	start := col
	for start > 0 && unicode.IsNumber(line[start-1]) {
		start--
	}

	end := col
	for end < len(line)-1 && unicode.IsNumber(line[end+1]) {
		end++
	}

	numBuffer := ""
	for i := start; i <= end; i++ {
		numBuffer += string(schematic[row][i])
	}

	number, err := strconv.Atoi(numBuffer)
	if err != nil {
		return 0, "", false
	}

	position := fmt.Sprintf("%d-%d", start, end)

	return number, position, true
}
