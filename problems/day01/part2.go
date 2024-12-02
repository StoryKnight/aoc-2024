package day01

import "fmt"

func Part2() {

	similarity := 0
	left, right := extract()
	// get a hashmap of each one
	// leftMap := make(map[int]int)
	rightMap := make(map[int]int)

	for _, value := range right {

		if _, exists := rightMap[value]; !exists {
			rightMap[value] = 1
		} else {
			rightMap[value] = rightMap[value] + 1
		}
	}

	for _, value := range left {
		similarity += value * rightMap[value]
	}

	fmt.Printf("similarity score: %d\n", similarity)

}
