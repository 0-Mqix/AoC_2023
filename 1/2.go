//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var convert = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

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

		for i, char := range line {

			if unicode.IsDigit(char) {
				fmt.Println("digit", string(char))

				if first == '?' {
					first = char
				}

				last = char

				continue
			}

			for word, c := range convert {
				if i+len(word)-1 > len(line)-1 {
					continue
				}

				selection := line[i : i+len(word)]

				if selection == word {
					fmt.Println("word", string(c))

					if first == '?' {
						first = c
					}
					last = c

					continue
				}
			}
		}

		fmt.Println(string(first), string(last))

		x, _ := strconv.Atoi(string(first) + string(last))
		total += x
	}

	fmt.Println(total)
}
