package main

import (
	"fmt"
	"strconv"
)

func mapRaces(lines *[]string) [][]int {
	times := []int{}
	distances := []int{}
	races := [][]int{}

	l := 4
	r := 4
	for r < len((*lines)[0]) {
		if string((*lines)[0][r]) == " " || r == len((*lines)[0])-1 {
			subStr := ""
			if r == len((*lines)[0])-1 {
				subStr = (*lines)[0][l+1 : r+1]
			} else {
				subStr = (*lines)[0][l+1 : r]
			}
			num, err := strconv.Atoi(string(subStr))
			if err == nil {
				times = append(times, num)
			}
			l = r
		}
		r++
	}

	l = 8
	r = 8
	for r < len((*lines)[1]) {
		if string((*lines)[1][r]) == " " || r == len((*lines)[1])-1 {
			subStr := ""
			if r == len((*lines)[1])-1 {
				subStr = (*lines)[1][l+1 : r+1]
			} else {
				subStr = (*lines)[1][l+1 : r]
			}
			num, err := strconv.Atoi(string(subStr))
			if err == nil {
				distances = append(distances, num)
			}
			l = r
		}
		r++
	}

	for i := 0; i < len(times); i++ {
		races = append(races, []int{times[i], distances[i]})
	}

	return races
}
func partOne(lines *[]string) {
	races := mapRaces(lines)

	// time = 7; distance = 9
	// charge for 7, have 0 seconds to go 0 distance
	// charge for 6, have 1 second to go 6 distance
	// charge for 5, have 2 seconds to go 10 distance
	// charge for 4, have 3 seconds to go 12 distance
	// charge for 3, have 4 seconds to go 12 distance
	// charge for 2, have 5 seconds to go 10 distance
	// charge for 1, have 7 seconds to go 7 distance

	// (time - charge time) * charge time >= distance
	// (7 - 4) * 4 >= 9
	// 12 >= 9

	// could just find the upper bound and lower bound, knowing that everything between is valid

	sum := 0
	for r := range races {
		lowerBound := 0
		flag := false
		charge := 1
		for !flag {
			if (races[r][0]-charge)*charge > races[r][1] {
				lowerBound = charge
				flag = true
			} else {
				charge += 1
			}
		}

		upperBound := 0
		flag = false
		charge = races[r][0] - 1
		for !flag {
			if (races[r][0]-charge)*charge > races[r][1] {
				upperBound = charge
				flag = true
			} else {
				charge -= 1
			}
		}

		if sum == 0 {
			sum = upperBound - lowerBound + 1
		} else {
			sum = sum * (upperBound - lowerBound + 1)
		}
	}

	fmt.Println(sum)
}
