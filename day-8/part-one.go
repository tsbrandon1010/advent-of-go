package main

import "fmt"

func partOne(lines *[]string) {

	instructions := (*lines)[0]
	mapping := map[string][]string{}
	for i := 2; i < len((*lines)); i++ {
		mapping[(*lines)[i][:3]] = []string{(*lines)[i][7:10], (*lines)[i][12:15]}
	}

	steps := 0
	currStep := "AAA"
	instructionPointer := 0
	for currStep != "ZZZ" {
		steps++

		if instructions[instructionPointer] == 'L' {
			currStep = mapping[currStep][0]
		} else {
			currStep = mapping[currStep][1]
		}

		instructionPointer++
		if instructionPointer >= len(instructions) {
			instructionPointer = 0
		}

	}

	fmt.Println(steps)
}
