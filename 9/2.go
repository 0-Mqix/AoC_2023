//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var example = `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func containsZero(row []int) bool {
	for _, x := range row {
		if x != 0 {
			return false
		}
	}
	return true
}

func parseIntArray(s string) []int {
	array := make([]int, 0)

	for _, s := range strings.Split(s, " ") {

		if len(s) == 0 {
			continue
		}

		x, _ := strconv.Atoi(s)
		array = append(array, x)
	}

	return array
}

func calculate(row []int) int {
	lastRow := row
	firsts := make([]int, 0)

	for !containsZero(row) {
		size := len(row)
		firsts = append(firsts, row[0])
		lastRow = row

		row = make([]int, size-1)

		for i := size - 1; i > 0; i-- {
			row[i-1] = lastRow[i] - lastRow[i-1]
		}
	}

	result := firsts[len(firsts)-2] - lastRow[0]

	for i := len(firsts) - 2; i > 0; i-- {
		result = firsts[i-1] - result
	}

	return result
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

	result := 0
	for _, line := range strings.Split(input, "\n") {

		if len(line) == 0 {
			continue
		}

		result += calculate(parseIntArray(line))
	}

	fmt.Println(result)
}
