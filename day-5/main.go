package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
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

	partOne(&lines)

}

// an array of arrays that contain an array of type int
/*
	[
		[[1,2,3], [1,2,3]]
		[]
		[]
	]

*/
