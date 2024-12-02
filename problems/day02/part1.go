package day02

import (
	"aoc-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	total := 0
	lines := utils.ReadLines("problems/day02/input.txt")

	for _, line := range lines {
		split := strings.Split(line, " ")
		// +/- 3 only, and only one direction
		valid := true

		prevTrend := 0 // 0 means no trend established yet
		prev, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}

		for _, c := range split[1:] {

			//fmt.Printf("prev: %d\n", prev)
			num, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println("Error converting string to integer: ", err)
				return
			}

			//fmt.Printf("c: %d\n", num)

			trend := 0
			if num > prev {
				trend = 1
			} else if num < prev {
				trend = -1
			} else {
				// equal
				valid = false
			}

			if trend != 0 && trend != prevTrend {
				if prevTrend != 0 {
					//fmt.Println("trend broken!")
					valid = false

				}
				prevTrend = trend // Update the trend
			}

			x := num - prev
			if x == 0 {
				valid = false
			}
			if x > 3 || x < -3 {
				valid = false
			}

			prev = num
		}

		if valid {
			total += 1
		}

	}

	fmt.Printf("total valid: %d\n", total)
}
