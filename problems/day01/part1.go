package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diff(num1 int, num2 int) int {
	res := num2 - num1
	if res < 0 {
		res = -res
	}
	return res
}

func extract() ([]int, []int) {

	file, err := os.Open("problems/day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Algorithm:
	// nlog(n)

	// sort right and left columns
	// get diff between columns and iterate, add to total
	// print total

	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text() // first, let's get the input
		split := strings.Split(line, "   ")

		// convert first integer
		num1, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return nil, nil
		}

		// convert second integer
		num2, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return nil, nil
		}

		left = append(left, num1)
		right = append(right, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return left, right
}

func Part1() {

	left, right := extract()

	// now sort both columns
	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i, value := range left {
		total += diff(value, right[i])
	}

	fmt.Printf("total: %d\n", total)
}
