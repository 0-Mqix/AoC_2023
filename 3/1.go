//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var example = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

var grid = [][]rune{}

func check(y, x int) bool {
	for yo := y - 1; yo < y+2; yo++ {

		if yo < 0 || yo >= len(grid) {
			continue
		}

		for xo := x - 1; xo < x+2; xo++ {

			if xo < 0 || xo >= len(grid[y]) {
				continue
			}

			char := grid[yo][xo]

			if !unicode.IsDigit(char) && char != '.' {
				return true
			}
		}
	}

	return false
}

func main() {
	input := example

	if len(os.Args) > 1 && os.Args[1] == "input" {
		raw, err := os.ReadFile("input")

		if err != nil {
			panic(err)
		}

		input = string(raw)
	}

	total := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune(line))
	}

	for y, line := range grid {
		digits := false
		adjacent := false
		number := ""

		for x, char := range line {

			if unicode.IsDigit(char) {
				if !adjacent {
					adjacent = check(y, x)
				}

				if !digits {
					digits = true
				}

				number += string(char)
			}

			if (!unicode.IsDigit(char) && digits) || x == len(line)-1 {
				if adjacent {
					x, _ := strconv.Atoi(number)
					total += x
				}

				number = ""
				digits = false
				adjacent = false
			}
		}
	}

	fmt.Println(total)
}
