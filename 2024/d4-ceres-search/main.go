package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Pos struct {
	x int
	y int
}

//go:embed input.txt
var input string

func main() {
	in := strings.Split(input, "\n")
	res := calc(in)
	fmt.Println(res)
}

func calc(in []string) int {
	var data [][]string
	for _, line := range in {
		x := strings.Split(line, "")
		data = append(data, x)
	}

	sum := 0

	directions := []Pos{
		{0, 1},
		{1, 1},
		{1, 0},
		{0, -1},
		{1, -1},
		{-1, 1},
		{-1, -1},
		{-1, 0},
	}

	for i, line := range data {
		for j, v := range line {
			if v != "X" {
				continue
			}

			pos := Pos{
				x: i,
				y: j,
			}
			for _, dir := range directions {
				if matchXMAS(pos, data, dir) {
					sum++
				}
			}
		}
	}
	return sum
}

func matchXMAS(pos Pos, data [][]string, dir Pos) bool {
	word := "XMAS"
	for i := 0; i < len(word); i++ {
		x := pos.x + dir.x*i
		y := pos.y + dir.y*i

		if x < 0 || x >= len(data) || y < 0 || y >= len(data[i]) {
			return false
		}

		if data[x][y] != string(word[i]) {
			return false
		}
	}

	return true
}
