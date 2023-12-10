package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	sum := 0

	cardCounts := map[int]int{}
	for card, line := range lines {
		cardCounts[card+1] += 1

		winningNums := map[string]int{}
		storedStr := ""
		for i := 42; i < len(line); i++ {
			if string(line[i]) == " " && storedStr != "" {
				winningNums[storedStr], _ = strconv.Atoi(storedStr)
				storedStr = ""
			} else {
				if string(line[i]) == " " {
					continue
				}
				storedStr += string(line[i])
			}
		}
		if storedStr != "" {
			winningNums[storedStr], _ = strconv.Atoi(storedStr)
		}

		currNum := ""
		runningSum := 0
		for i := 9; i < 40; i++ {
			if string(line[i]) == " " && currNum != "" {
				_, ok := winningNums[currNum]
				if ok {
					runningSum += 1
				}
				currNum = ""
			} else {
				if string(line[i]) == " " {
					continue
				}
				currNum += string(line[i])
			}
		}

		fmt.Print(card+1, runningSum, "- ")
		for i := card + 2; i < runningSum+card+2; i++ {
			cardCounts[i] += 1 * cardCounts[card+1]
			fmt.Print(i, ": ", cardCounts[i], " ")
		}
		fmt.Print("\n")
	}

	for _, value := range cardCounts {
		sum += value
	}

	fmt.Print(sum)

}
