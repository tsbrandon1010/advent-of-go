package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

type Hand struct {
	Cards    string
	Bid      string
	Type     int
	HighCard int
	Counts   map[string]int
}

var Ranks = map[string]int{
	"high":     1,
	"one-pair": 2,
	"two-pair": 3,
	"three":    4,
	"house":    5,
	"four":     6,
	"five":     7,
}

func countCards(hand *Hand) {
	for _, card := range hand.Cards {
		hand.Counts[string(card)] += 1
	}
}

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	partOne(&lines)
	partTwo(&lines)
}
