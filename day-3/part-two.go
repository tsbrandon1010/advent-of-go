package main

import (
	"fmt"
	"strconv"
)

func partTwo(grid *[140][140]string) {

	sum := 0
	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			if grid[i][j] == "*" {
				numOne := 0

				// check top
				if i-1 >= 0 {

					// check top left
					if j-1 >= 0 {
						num, err := strconv.Atoi(grid[i-1][j-1])
						if err == nil {
							if numOne == 0 {
								numOne = num
								fmt.Println(numOne)
							} else if numOne != 0 && num != numOne {
								fmt.Println(numOne, num)
								sum += numOne * num
								continue
							}
						}
					}

					// check top right
					if j+1 <= 140 {
						num, err := strconv.Atoi(grid[i-1][j+1])
						if err == nil {
							if numOne == 0 {
								numOne = num
								fmt.Println(numOne)

							} else if numOne != 0 && num != numOne {
								fmt.Println(numOne, num)

								sum += numOne * num
								continue
							}
						}
					}

					// check top
					num, err := strconv.Atoi(grid[i-1][j])
					if err == nil {
						if numOne == 0 {
							numOne = num
							fmt.Println(numOne)

						} else if numOne != 0 && num != numOne {
							fmt.Println(numOne, num)

							sum += numOne * num
							continue
						}
					}
				}

				// check bottom
				if i+1 < 140 {

					// check bottom left
					if j-1 >= 0 {
						num, err := strconv.Atoi(grid[i+1][j-1])
						if err == nil {
							if numOne == 0 {
								numOne = num
								fmt.Println(numOne)

							} else if numOne != 0 && num != numOne {
								fmt.Println(numOne, num)

								sum += numOne * num
								continue
							}
						}
					}
					// check bottom right
					if j+1 >= 0 {
						num, err := strconv.Atoi(grid[i+1][j+1])
						if err == nil {
							if numOne == 0 {
								numOne = num
								fmt.Println(numOne)

							} else if numOne != 0 && num != numOne {
								fmt.Println(numOne, num)

								sum += numOne * num
								continue
							}
						}
					}
					// check bottom
					num, err := strconv.Atoi(grid[i+1][j])
					if err == nil {
						if numOne == 0 {
							numOne = num
							fmt.Println(numOne)

						} else if numOne != 0 && num != numOne {
							fmt.Println(numOne, num)

							sum += numOne * num
							continue
						}
					}
				}

				// check left
				if j-1 >= 0 {
					num, err := strconv.Atoi(grid[i][j-1])
					if err == nil {
						if numOne == 0 {
							numOne = num
							fmt.Println(numOne)

						} else if numOne != 0 && num != numOne {
							fmt.Println(numOne, num)

							sum += numOne * num
							continue
						}
					}
				}

				// check right
				if j+1 >= 0 {
					num, err := strconv.Atoi(grid[i][j+1])
					if err == nil {
						if numOne != 0 && num != numOne {
							fmt.Println(numOne, num)

							sum += numOne * num
							continue
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
