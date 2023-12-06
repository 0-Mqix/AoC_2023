//go:build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var example = `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

var (
	mapRegex  = regexp.MustCompile(`(?ms)(\w+)-to-(\w+) map:\n((?:\d+ \d+ \d+\n?)+)`)
	seedRegex = regexp.MustCompile(`seeds: ((?:\d+(?: |))+)\n`)
)

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

func main() {
	input := example

	if len(os.Args) > 1 && os.Args[1] == "input" {
		raw, err := os.ReadFile("input")

		if err != nil {
			panic(err)
		}

		input = string(raw)
	}

	seeds := parseIntArray(seedRegex.FindStringSubmatch(input)[1])
	mappers := make([]func(int) int, 0)

	for _, match := range mapRegex.FindAllStringSubmatch(input, -1) {
		mappings := make([][]int, 0)

		for _, line := range strings.Split(match[3], "\n") {
			if len(line) == 0 {
				continue
			}
			mappings = append(mappings, parseIntArray(line))
		}

		mapper := func(input int) int {
			for _, array := range mappings {
				destinationStart := array[0]
				sourceStart := array[1]
				offset := destinationStart - sourceStart
				width := array[2]

				if input >= sourceStart && input < sourceStart+width {
					return input + offset
				}
			}

			return input
		}

		mappers = append(mappers, mapper)
	}

	var group sync.WaitGroup
	var mutex sync.Mutex
	var results = make([]int, 0)
	var count atomic.Int32

	full := len(seeds) / 2

	for i := 0; i < len(seeds); i += 2 {
		group.Add(1)
		go func(i int) {
			lowest := -1

			for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
				result := seed
				for _, mapper := range mappers {
					result = mapper(result)
				}

				if lowest == -1 || result < lowest {
					lowest = result
				}
			}

			mutex.Lock()
			results = append(results, lowest)
			progress := count.Add(1)
			fmt.Println("progress:", progress, "/", full)
			mutex.Unlock()

			group.Done()
		}(i)
	}

	group.Wait()

	fmt.Println("output:", slices.Min(results))
}
