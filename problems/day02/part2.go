package day02

import (
	"aoc-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func checkValid(split []string, prevTrend int, prev int) bool {
	for _, c := range split[1:] {

		//fmt.Printf("prev: %d\n", prev)
		num, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return false
		}

		//fmt.Printf("c: %d\n", num)

		trend := 0
		if num > prev {
			trend = 1
		} else if num < prev {
			trend = -1
		} else {
			// equal
			return false
		}

		if trend != 0 && trend != prevTrend {
			if prevTrend != 0 {
				//fmt.Println("trend broken!")
				return false
			}
			prevTrend = trend // Update the trend
		}

		diff := num - prev
		if diff == 0 {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
		prev = num
	}
	return true
}

func Part2() {
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

		valid = checkValid(split, prevTrend, prev)

		if valid {
			total += 1
		} else {
			// recursively call checkValid with a missing index (skip one)
			// if it ever returns true, we ball

			for i := 0; i < len(split); i++ {
				prevTrend = 0

				// ???
				modified := make([]string, 0, len(split)-1)
				modified = append(modified, split[:i]...)
				modified = append(modified, split[i+1:]...)

				prev, err = strconv.Atoi(modified[0])
				if err != nil {
					fmt.Println("Error converting string to integer: ", err)
					return
				}

				// Check if the modified array is valid
				if checkValid(modified, prevTrend, prev) {
					total += 1
					break
				}
			}
		}
	}

	fmt.Printf("total valid: %d\n", total)
}
