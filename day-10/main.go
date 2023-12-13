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

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func findNextMove(currPos rune, source rune) (rune, rune) {
	switch source {
	case 'N':
		// characters: '|', 'L', 'J'
		switch currPos {
		case '|':
			return 'S', 'N'
			// go south
		case 'L':
			return 'E', 'W'
			// go east
		case 'J':
			return 'W', 'E'
			// go west
		}

	case 'S':
		// characters: '|', 'F', '7'
		switch currPos {
		case '|':
			return 'N', 'S'
			// go north
		case 'F':
			return 'E', 'W'
			// go east
		case '7':
			return 'W', 'E'
			// go west
		}

	case 'E':
		// characters: '-', 'L', 'F'
		switch currPos {
		case '-':
			return 'W', 'E'
			// go west
		case 'L':
			return 'N', 'S'
			// go north
		case 'F':
			return 'S', 'N'
			// go south
		}

	case 'W':
		// characters: '-', 'J', '7'
		switch currPos {
		case '-':
			return 'E', 'W'
			// go east
		case 'J':
			return 'N', 'S'
			// go north
		case '7':
			return 'S', 'N'
			// go south
		}
	}

	return ' ', ' '
}

func findStarts(start []int, lines *[]string) [][]int {
	// start = {line #, column #}
	pointerStarts := [][]int{}

	// check up
	if start[0]-1 >= 0 {
		// valid characters: '|', '7', 'F'

		upChar := (*lines)[start[0]-1][start[1]]
		if upChar == '|' || upChar == '7' || upChar == 'F' {
			pointerStarts = append(pointerStarts, []int{start[0] - 1, start[1]})
		}
	}
	//check down
	if start[0]+1 < len(*lines) {
		// valid characters: '|', 'L', 'J'

		downChar := (*lines)[start[0]+1][start[1]]
		if downChar == '|' || downChar == 'L' || downChar == 'J' {
			pointerStarts = append(pointerStarts, []int{start[0] + 1, start[1]})
		}
	}

	//check left
	if start[1]-1 >= 0 {
		// valid characters: '-', 'L', 'F'

		leftChar := (*lines)[start[0]][start[1]-1]
		if leftChar == '-' || leftChar == 'L' || leftChar == 'F' {
			pointerStarts = append(pointerStarts, []int{start[0], start[1] - 1})
		}
	}

	//check right
	if start[1]+1 < len((*lines)[start[0]]) {
		// valid characters: '-', 'J', '7'

		rightChar := (*lines)[start[0]][start[1]+1]
		if rightChar == '-' || rightChar == 'J' || rightChar == '7' {
			pointerStarts = append(pointerStarts, []int{start[0], start[1] + 1})
		}
	}

	return pointerStarts
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	start := []int{0, 0}
	for l := range lines {
		for c := range lines[l] {
			if lines[l][c] == 'S' {
				start[0] = l
				start[1] = c
			}
		}
	}

	starts := findStarts(start, &lines)
	p1 := starts[0]

	var direction rune
	if p1[0] != start[0] {
		if p1[0] > start[0] {
			direction = 'N'
		} else {
			direction = 'S'
		}
	} else if p1[1] != start[1] {
		if p1[1] > start[1] {
			direction = 'W'
		} else {
			direction = 'E'
		}
	}

	var next rune
	count := 0
	for rune(lines[p1[0]][p1[1]]) != 'S' {

		next, direction = findNextMove(rune(lines[p1[0]][p1[1]]), direction)

		if next == 'S' {
			p1[0] += 1
		} else if next == 'N' {
			p1[0] -= 1
		} else if next == 'E' {
			p1[1] += 1
		} else if next == 'W' {
			p1[1] -= 1
		}
		count++
	}

	fmt.Println((count + 1) / 2)
}
