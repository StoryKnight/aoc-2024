package day05

import (
	"aoc-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part2() {

	input := utils.ReadLines("problems/day05/input.txt")

	// Parse rules
	// I misread this problem as add up all numbers after sorting
	// Got humbled by chatgpt I admit
	rules := make(map[int][]int)
	manualStart := false
	updates := [][]int{}

	for _, line := range input {
		if line == "" {
			manualStart = true
			continue
		}

		if manualStart {
			pageNumbers := strings.Split(line, ",")
			update := []int{}
			for _, page := range pageNumbers {
				num, _ := strconv.Atoi(page)
				update = append(update, num)
			}
			updates = append(updates, update)
		} else {
			parts := strings.Split(line, "|")
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])
			rules[num1] = append(rules[num1], num2)
		}
	}

	total := 0
	for _, update := range updates {
		if isOrdered(update, rules) {
			continue
		}

		fixedUpdate := topologicalSort(update, rules)

		mid := fixedUpdate[len(fixedUpdate)/2]
		total += mid
	}

	fmt.Printf("Total: %d\n", total)
}

func isOrdered(update []int, rules map[int][]int) bool {
	position := make(map[int]int)
	for i, num := range update {
		position[num] = i
	}

	for key, values := range rules {
		if posKey, exists := position[key]; exists {
			for _, val := range values {
				if posVal, ok := position[val]; ok && posVal < posKey {
					return false
				}
			}
		}
	}

	return true
}

func topologicalSort(update []int, rules map[int][]int) []int {
	// Build graph and in-degrees
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, num := range update {
		graph[num] = []int{}
		inDegree[num] = 0
	}

	for key, values := range rules {
		for _, val := range values {
			if _, exists := graph[key]; exists && contains(update, val) {
				graph[key] = append(graph[key], val)
				inDegree[val]++
			}
		}
	}

	// Kahn's Algorithm
	queue := []int{}
	for num, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, num)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
