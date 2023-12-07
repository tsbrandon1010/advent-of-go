package main

import (
	"bufio"
	"fmt"
	"os"
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

type GameSet struct {
	redCount   int
	greenCount int
	blueCount  int
}

type GameState struct {
	sets []GameSet
}

func main() {

	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range lines {
		l := 0
		r := 0

		num := ""
		color := ""

		for i := 0; i < len(line); i++ {

			if l != 0 && r != 0 && num == "" {
				num = line[l:r]
			}

			if line[i] == ',' {
				color = line[r+1 : i]
				break
			}

			if line[i] == ':' && l == 0 {
				l = i + 2
				i = l
			}
			if line[i] == ' ' && r == 0 && l != 0 {

				r = i
			}

		}

		fmt.Println(num, color)

	}

}
