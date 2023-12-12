package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ValuesP2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

func classifyTypeP2(hand *Hand) *Hand {
	for char, count := range (*hand).Counts {
		if char == "J" {
			continue
		}
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

	jCount := (*hand).Counts["J"]
	if jCount > 0 {
		if (*hand).Type == 6 {
			(*hand).Type = 7
		} else if (*hand).Type == 4 {
			if jCount == 1 {
				(*hand).Type = 6
			} else {
				(*hand).Type = 7
			}
		} else if (*hand).Type == 3 {
			(*hand).Type = 5
		} else if (*hand).Type == 2 {
			if jCount == 3 {
				(*hand).Type = 7
			} else if jCount == 2 {
				(*hand).Type = 6
			} else {
				(*hand).Type = 4
			}
		} else {
			if jCount == 5 {
				(*hand).Type = 7
			} else if jCount == 4 {
				(*hand).Type = 7
			} else if jCount == 3 {
				(*hand).Type = 6
			} else if jCount == 2 {
				(*hand).Type = 4
			} else if jCount == 1 {
				(*hand).Type = 2
			}
		}
	} else if (*hand).Type == 0 {
		// find the high card
		findHigh(hand)
		(*hand).Type = 1
	}

	return hand
}

func determineSecondOrderP2(handOne *Hand, handTwo *Hand) int {

	for i := 0; i < len(handOne.Cards); i++ {
		if ValuesP2[string(handOne.Cards[i])] == ValuesP2[string(handTwo.Cards[i])] {
			continue
		} else if ValuesP2[string(handOne.Cards[i])] > ValuesP2[string(handTwo.Cards[i])] {
			return 0
		} else {
			return 1
		}
	}

	return 0
}

func sortHandsP2(hands *[]*Hand) {
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
				if determineSecondOrderP2((*hands)[j], (*hands)[j+1]) == 0 {
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

func partTwo(lines *[]string) {

	hands := []*Hand{}
	for _, line := range *lines {
		splitLine := strings.Split(line, " ")
		hands = append(hands, &Hand{Cards: splitLine[0], Bid: splitLine[1], Counts: map[string]int{}, Type: 0})
	}

	for _, hand := range hands {
		countCards(hand)
		classifyTypeP2(hand)
	}

	sortHandsP2(&hands)

	sum := 0
	for rank, hand := range hands {
		num, err := strconv.Atoi(hand.Bid)
		if err == nil {
			sum += num * (rank + 1)
		}
	}

	fmt.Println(sum)
}
