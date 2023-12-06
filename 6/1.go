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
Time:      7  15   30
Distance:  9  40  200
`

func distance(hold, race int) int {
	return hold * (race - hold)
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

	numbers := make([][]int, 0)

	for _, line := range strings.Split(input, "\n") {

		if len(line) == 0 {
			continue
		}

		selecting := false
		number := ""

		array := make([]int, 0)

		for i, char := range line {

			digit := unicode.IsDigit(char)

			if digit {
				if !selecting {
					selecting = true
				}

				number += string(char)
			}

			if (!digit || i == len(line)-1) && selecting {
				x, _ := strconv.Atoi(number)
				array = append(array, x)
				number = ""
				selecting = false
			}
		}

		numbers = append(numbers, array)
	}

	margin := 1

	for i := 0; i < len(numbers[0]); i++ {
		time := numbers[0][i]
		record := numbers[1][i]
		wins := 0

		for i := 0; i < time; i++ {
			if distance(i, time) <= record {
				continue
			}
			wins++
		}

		margin *= wins
	}

	fmt.Println(margin)
}
