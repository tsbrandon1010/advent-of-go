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

	// create a 2D array
	// if run into an int, we scan until we reach the end of that int in the line
	// next, we replace every element we scanned with the full number

	// loop through the created array
	// if any int elements touch a symbol (up, down, left, right, diagonal) add to part to sum

	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	grid := [140][140]string{}

	currNum := ""
	numStart := -1
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {

			_, err := strconv.Atoi(string(lines[i][j]))
			if err == nil {
				if numStart == -1 {
					numStart = j
				}

				currNum += string(lines[i][j])

				if j == len(lines[i])-1 {
					for ns := numStart; ns <= j; ns++ {
						grid[i][ns] = currNum
					}

					numStart = -1
					currNum = ""
					break
				}

			} else if currNum != "" {
				for ns := numStart; ns < j; ns++ {
					grid[i][ns] = currNum
				}

				numStart = -1
				currNum = ""
			}

			grid[i][j] = string(lines[i][j])
		}
	}

	partOne(&grid)
	partTwo(&grid)
}
