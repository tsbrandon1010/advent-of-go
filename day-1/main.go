package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

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

func findFirst(s string) int {

	l := 0
	r := 0

	for r < len(s) {
		num, err := strconv.Atoi(string(s[l]))
		if err == nil {
			return num
		} else {
			for key, value := range numbers {
				if key == s[l:r+1] {
					return value
				}
			}
		}

		if r-l+1 >= 5 {
			l++
			r = l + 2
		} else if r == len(s)-1 {
			l++
		} else {
			r++
		}
	}
	return 0
}

func findLast(s string) int {

	r := len(s) - 1
	l := len(s) - 1

	for l >= 0 {
		num, err := strconv.Atoi(string(s[r]))
		if err == nil {
			return num
		} else {
			for key, value := range numbers {
				if key == s[l:r+1] {
					return value
				}
			}
		}

		if r-l+1 >= 5 {
			r--
			l = r - 2
		} else if l == 0 {
			r--
		} else {
			l--
		}
	}

	return 0
}

func main() {

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var ans int

	for _, str := range lines {

		first_num := findFirst(str)
		last_num := findLast(str)
		ans = ans + (first_num*10 + last_num)
	}

	fmt.Println(ans)
}
