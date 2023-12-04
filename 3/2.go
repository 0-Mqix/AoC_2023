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
var (
	grid  = [][]rune{}
	gears = map[string][]int{}
)

func check(y, x int) (bool, rune, int, int) {
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
				return true, char, yo, xo
			}
		}
	}

	return false, '.', -1, -1
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

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune(line))
	}

	for y, line := range grid {
		var digits bool
		var adjacent bool
		var number string
		var symbol rune
		sy, sx := -1, -1

		for x, char := range line {

			if unicode.IsDigit(char) {
				if !adjacent {
					adjacent, symbol, sy, sx = check(y, x)
				}

				if !digits {
					digits = true
				}

				number += string(char)
			}

			if (!unicode.IsDigit(char) && digits) || x == len(line)-1 {
				if adjacent && symbol == '*' {
					value, _ := strconv.Atoi(number)
					array, ok := gears[fmt.Sprintf("%d,%d", sy, sx)]
					if !ok {
						array = make([]int, 0)
					}
					array = append(array, value)
					gears[fmt.Sprintf("%d,%d", sy, sx)] = array
				}

				number = ""
				symbol = '.'
				digits = false
				adjacent = false
				sy, sx = -1, -1
			}
		}
	}

	total := 0

	for _, array := range gears {
		if len(array) == 2 {
			total += array[0] * array[1]
		}
	}

	fmt.Println(total)
}
