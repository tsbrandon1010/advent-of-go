package main

import (
	"fmt"
	"strconv"
)

func partOne(grid *[140][140]string) {

	flag := false
	sum := 0
	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			if grid[i][j] == "." {
				flag = false
				continue
			} else {
				num, err := strconv.Atoi(grid[i][j])
				if err == nil {

					if j-1 >= 0 && grid[i][j] == grid[i][j-1] && flag {
						continue
					}

					// check up + diagonal
					if i-1 >= 0 {
						// top left
						if j-1 >= 0 {
							_, err := strconv.Atoi(grid[i-1][j-1])
							if grid[i-1][j-1] != "." && err != nil {
								sum += num
								flag = true
								for j < len(grid[i])-2 && grid[i][j] == grid[i][j+1] {
									j++
								}
								continue
							}
						}

						// top right
						if j+1 < 140 {
							_, err := strconv.Atoi(grid[i-1][j+1])
							if grid[i-1][j+1] != "." && err != nil {
								sum += num
								flag = true
								for j < len(grid[i])-2 && grid[i][j] == grid[i][j+1] {
									j++
								}
								continue
							}
						}

						// top
						_, err := strconv.Atoi(grid[i-1][j])
						if grid[i-1][j] != "." && err != nil {
							sum += num
							flag = true
							for j < len(grid[i])-2 && grid[i][j] == grid[i][j+1] {

								j++
							}
							continue
						}
					}
					// check down + diagonal
					if i+1 < 140 {
						// bottom left
						if j-1 >= 0 {
							_, err := strconv.Atoi(grid[i+1][j-1])
							if grid[i+1][j-1] != "." && err != nil {
								sum += num
								flag = true
								for j < len(grid[i])-2 && grid[i][j] == grid[i][j+1] {

									j++
								}
								continue
							}
						}

						// bottom right
						if j+1 < 140 {
							_, err := strconv.Atoi(grid[i+1][j+1])
							if grid[i+1][j+1] != "." && err != nil {
								sum += num
								flag = true
								for j < len(grid[i])-2 && grid[i][j] == grid[i][j+1] {

									j++
								}
								continue
							}
						}
						_, err := strconv.Atoi(grid[i+1][j])
						if grid[i+1][j] != "." && err != nil {
							sum += num
							flag = true
							for j < len(grid[i])-2 && grid[i][j] == grid[i][j+1] {

								j++
							}
							continue
						}

					}

					// check left
					if j-1 >= 0 {
						_, err := strconv.Atoi(grid[i][j-1])
						if grid[i][j-1] != "." && err != nil {
							sum += num
							flag = true
							for j < len(grid[i])-2 && grid[i][j] == grid[i][j+1] {
								j++
							}
							continue
						}
					}
					// check right
					if j < 139 {
						_, err := strconv.Atoi(grid[i][j+1])
						if grid[i][j+1] != "." && err != nil {
							sum += num
							flag = true
							continue
						}
					}

				}
			}
		}
	}

	fmt.Println(sum)
}
