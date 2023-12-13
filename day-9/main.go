package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
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

	return lines, scanner.Err()
}

func extractHistory(lines *[]string) [][]int {

	histories := [][]int{}
	for _, line := range *lines {
		splitLine := strings.Split(line, " ")
		history := []int{}
		for _, val := range splitLine {
			num, _ := strconv.Atoi(val)
			history = append(history, num)
		}

		histories = append(histories, history)
	}

	return histories
}

func partOne(lastVal int, nums []int) (int, []int) {
	// this is our base case
	flag := true
	for _, num := range nums {
		if num != 0 {
			flag = false
			break
		}
	}
	if flag {
		return lastVal, nums
	}

	difference := []int{}
	for i := 0; i < len(nums)-1; i++ {
		difference = append(difference, nums[i+1]-nums[i])
	}

	return partOne(lastVal+difference[len(difference)-1], difference)
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	histories := extractHistory(&lines)

	sum := 0
	for _, history := range histories {
		endVal := history[len(history)-1]
		next, _ := partOne(endVal, history)
		sum += next
	}
	fmt.Println(sum)

	histories = extractHistory(&lines)
	sum = 0
	for i := range histories {
		history := reverseArray(histories[i])
		endVal := history[len(history)-1]
		next, _ := partOne(endVal, history)
		sum += next
	}
	fmt.Println(sum)
}
