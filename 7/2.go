//go:build ignore

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var example = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

type Type int

const (
	HIGH_CARD Type = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type Hand struct {
	bid   int
	cards [5]int
}

type Hands []*Hand

func (h Hands) Len() int      { return len(h) }
func (h Hands) Swap(a, b int) { h[a], h[b] = h[b], h[a] }
func (h Hands) Less(a, b int) bool {

	for i := 0; i < 5; i++ {
		if h[a].cards[i] == h[b].cards[i] {
			continue
		}
		return h[a].cards[i] < h[b].cards[i]
	}

	return false
}

var types = [7]Hands{}

func add(_type Type, hand *Hand) {
	types[_type] = append(types[_type], hand)
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

	for _, line := range strings.Split(input, "\n") {

		if len(line) == 0 {
			continue
		}

		data := strings.Split(line, " ")
		counts := [13]int{}
		hand := &Hand{}

		hand.bid, _ = strconv.Atoi(data[1])

		for i, card := range data[0] {
			id := strength(card)
			hand.cards[i] = id
			counts[id]++
		}

		biggest := 0
		pair := false
		two := false

		pairId := -1
		biggestId := -1

		for id, count := range counts {

			if id == 0 {
				continue
			}

			if count == 2 {
				if pair == true {
					two = true
				}

				pair = true
				pairId = id
			}

			if count > biggest {
				biggest = count
				biggestId = id
			}
		}

		switch biggest + counts[0] {
		case 5:
			add(FIVE_OF_A_KIND, hand)
		case 4:
			add(FOUR_OF_A_KIND, hand)
		case 3:
			if pair && pairId != biggestId {
				add(FULL_HOUSE, hand)
				break
			}
			add(THREE_OF_A_KIND, hand)
		default:
			if two {
				add(TWO_PAIR, hand)
				break
			}
			if pair || biggest+counts[0] == 2 {
				add(ONE_PAIR, hand)
				break
			}
			add(HIGH_CARD, hand)
		}
	}

	rank := 1
	winnings := 0

	for _, _type := range types {

		if len(_type) == 0 {
			continue
		}

		sort.Sort(_type)

		for _, hand := range _type {
			winnings += hand.bid * rank
			rank++
		}
	}

	fmt.Println(winnings)
}

func strength(char rune) int {
	switch char {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'T':
		return 9
	case '9':
		return 8
	case '8':
		return 7
	case '7':
		return 6
	case '6':
		return 5
	case '5':
		return 4
	case '4':
		return 3
	case '3':
		return 2
	case '2':
		return 1
	case 'J':
		return 0
	default:
		return -1
	}
}
