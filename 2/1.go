//go:build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	RED   = 12
	GREEN = 13
	BLUE  = 14
)

var (
	gameRegex = regexp.MustCompile(`(?m)Game (\d+)`)
	cubeRegex = regexp.MustCompile(`(?m)(\d+) (blue|red|green)(,|;)*`)
)

func main() {
	raw, err := os.ReadFile("input")

	if err != nil {
		panic(err)
	}

	input := string(raw)

	total := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		game := gameRegex.FindStringSubmatch(line)[1]
		id, _ := strconv.Atoi(game)

		data := strings.Split(line, ":")[1]

		valid := true
		currentRed := 0
		currentGreen := 0
		currentBlue := 0

		for _, v := range cubeRegex.FindAllStringSubmatch(data, -1) {
			amount, _ := strconv.Atoi(v[1])

			switch v[2] {
			case "red":
				currentRed += amount
			case "green":
				currentGreen += amount
			case "blue":
				currentBlue += amount
			}

			if currentRed > RED || currentGreen > GREEN || currentBlue > BLUE {
				valid = false
			}

			if v[3] == ";" {
				currentRed = 0
				currentGreen = 0
				currentBlue = 0
			}
		}

		if valid {
			total += id
		}
	}

	fmt.Println(total)
}
