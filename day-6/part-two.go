package main

import (
	"fmt"
	"strconv"
)

func mapRace(lines *[]string) []int {
	time := ""
	distance := ""
	race := []int{}

	r := 5
	for r < len((*lines)[0]) {
		if string((*lines)[0][r]) != " " {
			time += string((*lines)[0][r])
		}
		r++
	}

	r = 9
	for r < len((*lines)[1]) {
		if string((*lines)[1][r]) != " " {
			distance += string((*lines)[1][r])
		}
		r++
	}


	numTime, err := strconv.Atoi(time)
	if err == nil {
		race = append(race, numTime)
	}
	numDis, err := strconv.Atoi(distance)
	if err == nil {
		race = append(race, numDis)
	}

	return race
}

func partTwo(lines *[]string) {

	race := mapRace(lines)


	lowerBound := 0
	flag := false
	charge := 1
	for !flag {
		if (race[0]-charge)*charge > race[1] {
			lowerBound = charge
			flag = true
		} else {
			charge += 1
		}
	}

	upperBound := 0
	flag = false
	charge = race[0] - 1
	for !flag {
		if (race[0]-charge)*charge > race[1] {
			upperBound = charge
			flag = true
		} else {
			charge -= 1
		}
	}

	fmt.Println(upperBound - lowerBound + 1)
}
