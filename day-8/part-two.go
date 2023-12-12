package main

import "fmt"

func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a int, b int) int {
	return (a * b) / GCD(a, b)
}

func LCMList(nums []int) int {
	if len(nums) == 2 {
		return LCM(nums[0], nums[1])
	}
	return LCM(nums[0], LCMList(nums[1:]))
}

func partTwo(lines *[]string) {

	instructions := (*lines)[0]
	mapping := map[string][]string{}

	starts := []string{}
	for i := 2; i < len((*lines)); i++ {
		mapping[(*lines)[i][:3]] = []string{(*lines)[i][7:10], (*lines)[i][12:15]}
		if (*lines)[i][2] == 'A' {
			starts = append(starts, (*lines)[i][:3])
		}
	}

	endSteps := []int{}
	for i := range starts {
		steps := 0
		instructionPointer := 0
		for {
			if instructions[instructionPointer] == 'L' {
				starts[i] = mapping[starts[i]][0]
			} else {
				starts[i] = mapping[starts[i]][1]
			}
			steps++
			instructionPointer++
			if instructionPointer == len(instructions) {
				instructionPointer = 0
			}

			if starts[i][2] == 'Z' {
				endSteps = append(endSteps, steps)
				break
			}
		}

	}

	fmt.Println(LCMList(endSteps))
}
