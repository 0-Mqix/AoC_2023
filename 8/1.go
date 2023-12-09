//go:build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var example = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

var regex = regexp.MustCompile(`(?m)(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)

var network = map[string][2]string{}

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
	instructions := lines[0]

	for _, line := range lines[2:] {

		if len(line) == 0 {
			continue
		}

		data := regex.FindStringSubmatch(line)

		if len(data) != 4 {
			fmt.Println(data)
			continue
		}

		network[data[1]] = [2]string{data[2], data[3]}
	}

	location := "AAA"
	index := 0
	steps := 0

	for location != "ZZZ" {
		instruction := instructions[index]

		if instruction == 'L' {
			location = network[location][0]
		} else {
			location = network[location][1]
		}

		index++

		if index == len(instructions) {
			index = 0
		}

		steps++
	}

	fmt.Println(steps)
}
