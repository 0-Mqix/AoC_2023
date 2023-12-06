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

		number := ""
		array := make([]int, 0)

		for i, char := range line {

			if unicode.IsDigit(char) {
				number += string(char)
			}

			if i == len(line)-1 {
				x, _ := strconv.Atoi(number)
				array = append(array, x)
				number = ""
			}
		}

		numbers = append(numbers, array)
	}

	time := numbers[0][0]
	record := numbers[1][0]
	margin := 0

	for i := 0; i < time; i++ {
		if distance(i, time) <= record {
			continue
		}
		margin++
	}

	fmt.Println(margin)
}
