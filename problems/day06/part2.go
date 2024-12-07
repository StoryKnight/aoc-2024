package day06

import (
	"aoc-2024/utils"
	"fmt"
)

func Part2() {

	// row, col
	directions = [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	visited = make(map[[2]int]int)

	input := utils.ReadLines("problems/day06/input.txt")
	numRows := len(input)
	numCols := len(input[0])

	position := [2]int{0, 0}

	grid = make([][]rune, numRows)
	for i := range grid {
		grid[i] = make([]rune, numCols)
	}

	for i, line := range input {
		for j, char := range line {
			if char == '^' {
				position[0] = i
				position[1] = j
				grid[i][j] = '.'
			} else {
				grid[i][j] = char
			}
		}
	}

	// place an obstacle once at every possible position for findExit except start position

	cycles := 0

	for i, row := range grid {
		for j, col := range row {
			temp := col
			if [2]int{i, j} != position {
				grid[i][j] = '#'
				if !findExit(position, 0) {
					cycles += 1
				}
				utils.Clear(visited)
			}
			grid[i][j] = temp
		}
	}
	fmt.Printf("cycles: %d\n", cycles)
}
