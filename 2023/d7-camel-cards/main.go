package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Hand struct {
	Cards string
	Bid   int
}

type HandCategorized map[Type][]Hand

type Type int

const (
	HighCard Type = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

func main() {
	in := strings.Split(input, "\n")
	res := calculateCamelCardsGameWinnings(in)
	fmt.Println(res)
}

func calculateCamelCardsGameWinnings(in []string) int {
	camelCardHands, err := getHandsAndBids(in)
	if err != nil {
		return 0
	}

	handsCategorized := make(HandCategorized)

	for _, v := range camelCardHands {
		handType := v.categorizeHand()
		handsCategorized[handType] = append(handsCategorized[handType], v)
	}

	for _, v := range handsCategorized {
		sortHands(v)
	}

	res := 0
	rank := 1
	for i := HighCard; i <= FiveOfKind; i++ {
		hands := handsCategorized[i]
		for _, hand := range hands {
			res += hand.Bid * rank
			rank++
		}
	}

	return res
}

func getHandsAndBids(input []string) ([]Hand, error) {
	var camelCardHands []Hand
	for _, line := range input {
		x := strings.Split(line, " ")
		if len(x) != 2 {
			return nil, fmt.Errorf("invalid hand")
		}

		bid, err := strconv.Atoi(strings.TrimSpace(x[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid hand")
		}

		camelCardHand := Hand{
			Cards: x[0],
			Bid:   bid,
		}

		camelCardHands = append(camelCardHands, camelCardHand)
	}

	return camelCardHands, nil
}

func (hand Hand) categorizeHand() Type {
	handMap := make(map[rune]int)
	for _, v := range hand.Cards {
		handMap[v]++
	}

	pairs := 0
	hasThree := false

	for _, v := range handMap {
		if v == 2 {
			pairs++
		}

		if v == 3 {
			hasThree = true
		}

		if v == 4 {
			return FourOfKind
		}

		if v == 5 {
			return FiveOfKind
		}
	}

	if pairs == 1 && hasThree {
		return FullHouse
	}

	if hasThree {
		return ThreeOfKind
	}

	if pairs == 2 {
		return TwoPair
	}

	if pairs == 1 {
		return OnePair
	}

	return HighCard
}

func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		hand1, hand2 := hands[i].Cards, hands[j].Cards

		return compareHands(hand1, hand2)
	})
}

func compareHands(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {
		value1 := getCardValue(hand1[i])
		value2 := getCardValue(hand2[i])

		if value1 != value2 {
			return value1 < value2
		}
	}

	return false
}

func getCardValue(card byte) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}
