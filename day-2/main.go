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

type GameMax struct {
	redMax   int
	greenMax int
	blueMax  int
}

func main() {

	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	totalSum := 0
	for _, line := range lines {
		l := 0
		r := 0

		var currMax GameMax
		currSubStr := ""
		currNum := 0
		for r < len(line) {
			// put the left pointer at the start of the real string
			if line[r] == ':' && l == 0 {
				l = r + 2
				r++

				// if the left value is set then we can read from the right pointer
			} else if l != 0 {

				if line[r] == ' ' {
					i, err := strconv.Atoi(currSubStr)
					if err == nil {
						currNum = i
						currSubStr = ""
					}

				} else if line[r] == ',' || line[r] == ';' || r == len(line)-1 {
					if r == len(line)-1 {
						currSubStr += string(line[r])
					}

					switch currSubStr {
					case "red":
						if currNum > currMax.redMax {
							currMax.redMax = currNum
						}
					case "green":
						if currNum > currMax.greenMax {
							currMax.greenMax = currNum
						}
					case "blue":
						if currNum > currMax.blueMax {
							currMax.blueMax = currNum
						}
					default:
					}

					currSubStr = ""
				} else {
					currSubStr += string(line[r])
				}
			}

			r++
		}

		power := currMax.redMax * currMax.greenMax * currMax.blueMax
		totalSum += power
		fmt.Println(totalSum)

	}

	fmt.Println(totalSum)
}
