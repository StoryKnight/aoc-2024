package day04

import (
	"aoc-2024/utils"
	"fmt"
)

func Part1() {
	lines := utils.ReadLines("problems/day04/input.txt")
	// input appears to be a perfect square

	total := 0
	numRows := len(lines)
	numCols := len(lines[0])

	// is there a better way than this . . .
	for i, row := range lines {
		for j, col := range row {
			if col == 'X' {
				// check up
				if i >= 3 {
					if lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
						total += 1
					}
				}

				// check down
				if i < numRows-3 {
					if lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
						total += 1
					}
				}

				// check left
				if j >= 3 {
					if lines[i][j-1] == 'M' && lines[i][j-2] == 'A' && lines[i][j-3] == 'S' {
						total += 1
					}
				}

				// check right
				if j < numCols-3 {
					if lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
						total += 1
					}
				}

				// check up-left
				if i >= 3 && j >= 3 {
					if lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
						total += 1
					}
				}

				// check up-right
				if i >= 3 && j < numCols-3 {
					if lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
						total += 1
					}
				}

				// check down-left
				if i < numRows-3 && j >= 3 {
					if lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
						total += 1
					}
				}

				// check down-right
				if i < numRows-3 && j < numCols-3 {
					if lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
						total += 1
					}
				}
			}
		}
	}

	fmt.Printf("total: %d\n", total)

}
