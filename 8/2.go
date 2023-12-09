//go:build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var example = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

var regex = regexp.MustCompile(`(?m)(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)

var (
	network      = map[string][2]string{}
	instructions = ""
)

// looked at reddit so i cheated a little bit (dit not know wat a LCM is or a CRT)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func walk(location string) int {
	index := 0
	steps := 0

	for location[2] != 'Z' {

		instruction := instructions[index]

		if instruction == 'L' {
			location = network[location][0]
		} else {
			location = network[location][1]
		}

		index++
		steps++

		if index == len(instructions) {
			index = 0
		}
	}

	return steps
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
	instructions = lines[0]

	for _, line := range lines[2:] {

		if len(line) == 0 {
			continue
		}

		data := regex.FindStringSubmatch(line)

		if len(data) != 4 {
			continue
		}

		network[data[1]] = [2]string{data[2], data[3]}
	}

	result := 0

	for location := range network {

		if location[2] != 'A' {
			continue
		}

		if result == 0 {
			result = walk(location)
		} else {
			result = lcm(result, walk(location))
		}

	}

	fmt.Println(result)
}
