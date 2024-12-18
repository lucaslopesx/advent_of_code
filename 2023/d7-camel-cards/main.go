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

	joker := 0
	for _, v := range hand.Cards {
		if v == 'J' {
			joker++
			continue
		}
		handMap[v]++
	}

	if joker == 5 {
		return FiveOfKind
	}

	max := 0
	for _, count := range handMap {
		if count > max {
			max = count
		}
	}

	max += joker
	if max == 5 {
		return FiveOfKind
	}

	if max == 4 {
		return FourOfKind
	}

	pairs := 0
	for _, count := range handMap {
		if count == 2 {
			pairs++
		}
	}

	if max == 3 {
		if pairs == 2 || (pairs == 1 && joker == 0) {
			return FullHouse
		}

		return ThreeOfKind
	}

	if pairs == 2 {
		return TwoPair
	}

	if max == 2 {
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
		return 1
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}
