//go:build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var example = `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

var (
	regex = regexp.MustCompile(`(?m)Card +(\d+): +(.+) +\| +(.+)`)
)

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

	for _, line := range strings.Split(input, "\n") {
		lineTotal := 0

		if len(line) == 0 {
			continue
		}

		match := regex.FindStringSubmatch(line)
		winning := make(map[string]bool)
		won := make([]string, 0)

		for _, x := range strings.Split(match[2], " ") {
			x = strings.TrimSpace(x)

			if len(x) == 0 {
				continue
			}

			winning[x] = true
		}

		for _, x := range strings.Split(match[3], " ") {
			x = strings.TrimSpace(x)

			if _, win := winning[x]; win {

				won = append(won, x)

				if lineTotal == 0 {
					lineTotal = 1
					continue
				}

				lineTotal *= 2
			}
		}

		total += lineTotal
	}

	fmt.Println(total)
}
