package day4

import (
	"fmt"
)

const roll = byte('@')
const empty = byte('.')

func isOutBounds(i, j, rows, cols int) bool {
	if i < 0 || i >= rows {
		return true
	}
	if j < 0 || j >= cols {
		return true
	}
	return false
}

func Solution(input [][]byte) int {

	directions := [8][2]int{
		{1, 1},
		{1, 0},
		{1, -1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{0, 1},
	}

	count := 0
	rows, cols := len(input), len(input[0])
	for i := range rows {
		for j := range cols {
			if input[i][j] == empty {
				continue
			}
			rolls := 0
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if isOutBounds(i+dx, j+dy, rows, cols) {
					continue
				}
				if input[i+dx][j+dy] == roll {
					rolls += 1
				}
			}
			if rolls < 4 {
				fmt.Println("found here", i, j, rolls)
				count += 1
			}
		}
	}
	return count
}
