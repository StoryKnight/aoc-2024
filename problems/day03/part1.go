package day03

import (
	"aoc-2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func matchMul() [][]string {

	input, err := utils.ReadFileToString("problems/day03/input.txt")
	if err != nil {
		fmt.Println("cannot read file")
	}

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	return re.FindAllStringSubmatch(input, -1)
}

func Part1() {

	sum := 0
	matches := matchMul()

	if matches == nil {
		fmt.Println("no match")
		return
	}

	for _, num := range matches {
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

		sum += x * y
	}

	fmt.Printf("sum: %d", sum)

}
