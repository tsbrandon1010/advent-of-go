package main

import (
	"fmt"
	"strconv"
	"strings"
)

func loadSeeds(line string) []int {
	seeds := []int{}
	l := 5
	r := 5
	for r < len(line) {
		if string(line[r]) == " " || r == len(line)-1 {
			subStr := ""
			if r == len(line)-1 {
				subStr = line[l+1 : r+1]
			} else {
				subStr = line[l+1 : r]
			}
			num, err := strconv.Atoi(string(subStr))
			if err == nil {
				seeds = append(seeds, num)
			}
			l = r
		}
		r++
	}
	return seeds
}
func partOne(lines *[]string) {

	seeds := loadSeeds((*lines)[0])

	mappings := [][][]int{}
	for _, line := range (*lines)[2:] {
		if line == "" {
			continue
		}
		if line[len(line)-4:] == "map:" {
			mappings = append(mappings, [][]int{})
		} else {

			splitString := strings.Split(line, " ")
			val1, _ := strconv.Atoi(splitString[0])
			val2, _ := strconv.Atoi(splitString[1])
			val3, _ := strconv.Atoi(splitString[2])

			mappings[len(mappings)-1] = append(mappings[len(mappings)-1], []int{val1, val2, val3})
		}
	}

	seedPath := map[int]int{}
	for _, seed := range seeds {
		seedPath[seed] = seed

		for _, currMap := range mappings {
			for _, vals := range currMap {

				if vals[1] <= seedPath[seed] && seedPath[seed] < vals[1]+vals[2] {
					seedPath[seed] = vals[0] + (seedPath[seed] - vals[1])
					break
				}
			}
		}
	}

	min := seedPath[seeds[0]]
	for _, val := range seedPath {
		if val < min {
			min = val
		}
	}

	fmt.Print(min)
}
