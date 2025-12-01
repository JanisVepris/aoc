package day07

import (
	"fmt"
	"sort"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var (
	lines       []string
	jackIsJoker = false
)

func Setup() {
	lines = files.ReadFile("2023/day07/input.txt")
}

func Part1() {
	result := 0

	hands := sortHands(buildHands(lines))

	for rank, hand := range hands {
		result += hand.bid * (rank + 1)
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	jackIsJoker = true

	result := 0

	hands := sortHands(buildHands(lines))

	for rank, hand := range hands {
		result += hand.bid * (rank + 1)
	}
	fmt.Printf("Part 2: %d\n", result)
}

type HandValue int

const (
	HighCard     HandValue = iota
	Pair         HandValue = iota
	TwoPair      HandValue = iota
	ThreeOfAKind HandValue = iota
	FullHouse    HandValue = iota
	FourOfAKind  HandValue = iota
	FiveOfAKind  HandValue = iota
)

type Card struct {
	symbol string
}

func (c *Card) Value() int {
	switch c.symbol {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if jackIsJoker {
			return 1
		}
		return 11
	case "T":
		return 10
	default:
		return conv.StrToInt(c.symbol)
	}
}

type Hand struct {
	cards []Card
	value HandValue
	bid   int
}

func (h *Hand) Card(i int) Card {
	return h.cards[i]
}

func sortHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {

			for k := 0; k < len(hands[i].cards); k++ {
				if hands[i].cards[k].Value() == hands[j].cards[k].Value() {
					continue
				}

				return hands[i].cards[k].Value() < hands[j].cards[k].Value()
			}

			return false
		}

		return hands[i].value < hands[j].value
	})

	return hands
}

func buildHands(lines []string) []Hand {
	hands := make([]Hand, 0)

	for _, line := range lines {
		parts := array.Filter(strings.Split(line, " "), func(s string) bool { return s != "" })

		cards, handValue := parseCards(parts[0])

		hand := Hand{cards: cards, bid: conv.StrToInt(parts[1]), value: handValue}

		hands = append(hands, hand)
	}

	return hands
}

func parseCards(cardString string) ([]Card, HandValue) {
	cards := make([]Card, 0)

	for _, c := range cardString {
		cards = append(cards, Card{symbol: string(c)})
	}

	return cards, getHandValue(cards)
}

func getHandValue(cards []Card) HandValue {
	hasPair := false
	hasThreeOfAKind := false

	cardCounts := make(map[string]int)

	for _, card := range cards {
		cardCounts[card.symbol]++
	}

	if jackIsJoker && cardCounts["J"] > 0 && cardCounts["J"] != 5 {
		maxCount := 0
		maxCountSymbol := ""

		for cardSymbol, count := range cardCounts {
			if cardSymbol == "J" {
				continue
			}

			if count > maxCount {
				maxCount = count
				maxCountSymbol = cardSymbol
			}

		}

		cardCounts[maxCountSymbol] += cardCounts["J"]
		cardCounts["J"] = 0
	}

	for _, count := range cardCounts {
		if count == 0 {
			continue
		} else if count == 5 {
			return FiveOfAKind
		} else if count == 4 {
			return FourOfAKind
		} else if count == 3 {
			hasThreeOfAKind = true
		} else if count == 2 {
			if hasPair {
				return TwoPair
			} else {
				hasPair = true
			}
		}
	}

	if hasThreeOfAKind && hasPair {
		return FullHouse
	}
	if hasThreeOfAKind {
		return ThreeOfAKind
	}
	if hasPair {
		return Pair
	}

	return HighCard
}
