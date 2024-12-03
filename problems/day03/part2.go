package day03

import (
	"aoc-2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func matchDoMul() [][]string {

	input, err := utils.ReadFileToString("problems/day03/input.txt")
	if err != nil {
		fmt.Println("cannot read file")
	}

	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	return re.FindAllStringSubmatch(input, -1)
}

func Part2() {
	sum := 0
	matches := matchDoMul()

	if matches == nil {
		fmt.Println("no match")
		return
	}

	do := true
	for _, num := range matches {

		if num[0] == "don't()" {
			do = false
			continue
		} else if num[0] == "do()" {
			do = true
			continue
		}

		x, err := strconv.Atoi(num[1])
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}

		y, err := strconv.Atoi(num[2])
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}

		if do {
			sum += x * y
		}
	}

	fmt.Printf("sum: %d", sum)

}
