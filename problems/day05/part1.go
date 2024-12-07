package day05

import (
	"aoc-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	input := utils.ReadLines("problems/day05/input.txt")
	manual := false

	// iterate backward on manual
	// lookup via hashmap, get values, see if any previous number is on it
	// if it's on the list, it's bad (rule break)

	rules := make(map[int][]int)
	total := 0

	for _, line := range input {

		if line == "" {
			manual = true
		}

		if manual {
			// parse the manual
			valid := true
			order := strings.Split(line, ",")
			size := len(order) - 1
			middle := 0

			// iterate backward
			for i := range order {

				//fmt.Println(order)
				num, err := strconv.Atoi((order[size-i]))
				if err != nil {
					fmt.Println("error parsing int")
				}

				if i == size/2 {
					middle = num
				}

				for j := size - 1 - i; j >= 0; j-- {
					//fmt.Println("iterate backward")
					candidate, err := strconv.Atoi(order[j])
					if err != nil {
						fmt.Println("error parsing int")
					}
					//fmt.Printf("candidate: %d\n", candidate)

					for _, v := range rules[num] {
						if v == candidate {
							valid = false
						}
					}
				}
			}

			if valid {
				total += middle
			}

		} else {
			// parse the rules
			rule := strings.Split(line, "|")
			num1, err := strconv.Atoi((rule[0]))
			if err != nil {
				fmt.Println("error parsing int")
			}
			num2, err := strconv.Atoi((rule[1]))
			if err != nil {
				fmt.Println("error parsing int")
			}
			rules[num1] = append(rules[num1], num2)
		}
	}

	fmt.Printf("total: %d\n", total)

}
