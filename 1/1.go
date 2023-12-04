//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	raw, err := os.ReadFile("input")

	if err != nil {
		panic(err)
	}

	input := string(raw)

	var total int

	for _, line := range strings.Split(input, "\n") {
		first := '?'
		var last rune

		for _, char := range line {
			if !unicode.IsDigit(char) {
				continue
			}

			if first == '?' {
				first = char
			}

			last = char
		}

		x, _ := strconv.Atoi(string(first) + string(last))
		total += x
	}

	fmt.Println(total)
}
