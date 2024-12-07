package day06

import (
	"aoc-2024/utils"
	"fmt"
)

var grid [][]rune
var visited map[[2]int]int
var directions [][2]int

func drawGrid(position [2]int, direction int) {
	for i, row := range grid {
		for j, char := range row {
			if i == position[0] && j == position[1] {
				switch direction % 4 {
				case 0:
					fmt.Print("^ ")
				case 1:
					fmt.Print("> ")
				case 2:
					fmt.Print("v ")
				case 3:
					fmt.Print("< ")
				default:
					fmt.Print("? ")
				}
			} else {
				fmt.Print(string(char), " ")
			}
		}
		fmt.Println()
	}
}

// return true if there is an exit, eventually
// return false if you enter samsara
func findExit(position [2]int, direction int) bool {

	drawGrid(position, direction) // draw because fun

	if _, exists := visited[position]; !exists {
		visited[position] = direction % 4
	} else {
		// we've been here, check for cycle (this is a part2 addition)
		if visited[position] == direction%4 {
			return false
		}
	}

	candidateX := position[1] + directions[direction%4][1]
	candidateY := position[0] + directions[direction%4][0]

	// bounds-check
	if candidateX >= len(grid[0]) || candidateY >= len(grid) || candidateY < 0 || candidateX < 0 {
		return true
	}

	if grid[candidateY][candidateX] == '#' {
		direction += 1
		return findExit(position, direction)
	}

	// turn if facing obstacle
	position[1] = candidateX
	position[0] = candidateY

	return findExit(position, direction)
}

func Part1() {

	// row, col

	directions = [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
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
				fmt.Printf("we start at row(%d), col(%d)\n", i, j)
				position[0] = i
				position[1] = j
				grid[i][j] = '.'
			} else {
				grid[i][j] = char
			}
		}
	}

	findExit(position, 0)
	fmt.Printf("visited: %d", len(visited))
}
