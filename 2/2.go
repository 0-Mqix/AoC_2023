//go:build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var example = `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

var (
	gameRegex = regexp.MustCompile(`(?m)Game (\d+)`)
	cubeRegex = regexp.MustCompile(`(?m)(\d+) (blue|red|green)(,|;)*`)
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
		if len(line) == 0 {
			continue
		}

		data := strings.Split(line, ":")[1]

		valid := true
		currentRed := 0
		currentGreen := 0
		currentBlue := 0

		for _, v := range cubeRegex.FindAllStringSubmatch(data, -1) {
			amount, _ := strconv.Atoi(v[1])

			switch v[2] {
			case "red":
				if amount > currentRed {
					currentRed = amount
				}
			case "green":
				if amount > currentGreen {
					currentGreen = amount
				}
			case "blue":
				if amount > currentBlue {
					currentBlue = amount
				}
			}
		}

		if valid {
			total += currentRed * currentGreen * currentBlue
		}
	}

	fmt.Println(total)
}
