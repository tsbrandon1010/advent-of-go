package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ValuesP1 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func findHigh(hand *Hand) {
	hand.HighCard = ValuesP1[string(hand.Cards[0])]
	for _, card := range hand.Cards {
		if ValuesP1[string(card)] > hand.HighCard {
			hand.HighCard = ValuesP1[string(card)]
		}
	}
}

func classifyTypeP1(hand *Hand) *Hand {
	for _, count := range (*hand).Counts {
		if count == 2 {
			if (*hand).Type == 2 {
				(*hand).Type = 3
			} else if (*hand).Type == 4 {
				(*hand).Type = 5
			} else {
				(*hand).Type = 2
			}
		} else if count == 3 {
			if (*hand).Type == 2 {
				(*hand).Type = 5
			} else {
				(*hand).Type = 4
			}
		} else if count == 4 {
			(*hand).Type = 6
		} else if count == 5 {
			(*hand).Type = 7
		}
	}

	if (*hand).Type == 0 {
		// find the high card
		findHigh(hand)
		(*hand).Type = 1
	}

	return hand
}

func determineSecondOrderP1(handOne *Hand, handTwo *Hand) int {

	for i := 0; i < len(handOne.Cards); i++ {
		if ValuesP1[string(handOne.Cards[i])] == ValuesP1[string(handTwo.Cards[i])] {
			continue
		} else if ValuesP1[string(handOne.Cards[i])] > ValuesP1[string(handTwo.Cards[i])] {
			return 0
		} else {
			return 1
		}
	}

	return 0
}

func sortHandsP1(hands *[]*Hand) {
	n := len(*hands)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {

			if (*hands)[j].Type == 1 && (*hands)[j+1].Type == 1 && (*hands)[j].HighCard > (*hands)[j+1].HighCard {
				temp := (*hands)[j+1]
				(*hands)[j+1] = (*hands)[j]
				(*hands)[j] = temp
				swapped = true

			} else if (*hands)[j].Type > (*hands)[j+1].Type {
				temp := (*hands)[j+1]
				(*hands)[j+1] = (*hands)[j]
				(*hands)[j] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if (*hands)[j].Type == (*hands)[j+1].Type {
				if determineSecondOrderP1((*hands)[j], (*hands)[j+1]) == 0 {
					temp := (*hands)[j+1]
					(*hands)[j+1] = (*hands)[j]
					(*hands)[j] = temp
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
}

func partOne(lines *[]string) {

	// 1. classify hands
	// - count each character in hand
	// - number of duplicates determines hand type

	// 2. rank each hand
	// - First, rank based on type
	// - if there are duplicate types, rank based on second ordering

	// 3. determine winnings

	hands := []*Hand{}
	for _, line := range *lines {
		splitLine := strings.Split(line, " ")
		hands = append(hands, &Hand{Cards: splitLine[0], Bid: splitLine[1], Counts: map[string]int{}, Type: 0})

	}

	for _, hand := range hands {
		countCards(hand)
		classifyTypeP1(hand)
	}

	sortHandsP1(&hands)

	sum := 0
	for rank, hand := range hands {
		num, err := strconv.Atoi(hand.Bid)
		if err == nil {
			sum += num * (rank + 1)
		}
	}

	fmt.Println(sum)
}
