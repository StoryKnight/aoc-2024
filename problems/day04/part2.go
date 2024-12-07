package day04

import (
	"aoc-2024/utils"
	"fmt"
)

func Part2() {
	total := 0
	lines := utils.ReadLines("problems/day04/input.txt")

	numRows := len(lines)
	numCols := len(lines[0])

	for i, row := range lines {
		for j, col := range row {
			if col == 'A' {
				// check up-left -> down-right
				if i >= 1 && j >= 1 && i < numRows-1 && j < numCols-1 {
					// top-left to down-right
					tLdR := (lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')
					// top-right to down-left
					tRdL := (lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S') || (lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M')
					if tLdR && tRdL {
						total += 1
					}
				}
			}
		}
	}

	fmt.Printf("total: %d\n", total)
}
