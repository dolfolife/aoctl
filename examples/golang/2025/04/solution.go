package main

import (
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	matrix := make([][]int, 0)
	rows := strings.Split(input, "\n")
	result := 0

	for _, row := range rows {
		matrix = append(matrix, make([]int, 0))
		for _, char := range row {
			if char == '@' {
				matrix[len(matrix)-1] = append(matrix[len(matrix)-1], 1)
			} else {
				matrix[len(matrix)-1] = append(matrix[len(matrix)-1], 0)
			}
		}
	}

	for i, row := range matrix {
		for j, cell := range row {
			if cell == 0 {
				continue
			}
			if !hasMoreThanFourAdjacentPapers(i, j, matrix) {
				result++
			}
		}
	}

	return strconv.Itoa(result), nil
}

func hasMoreThanFourAdjacentPapers(i int, j int, matrix [][]int) bool {
	result := 0
	if i > 0 {
		result += matrix[i-1][j]
	}
	if i < len(matrix)-1 {
		result += matrix[i+1][j]
	}
	if j > 0 {
		result += matrix[i][j-1]
	}
	if j < len(matrix[i])-1 {
		result += matrix[i][j+1]
	}
	if i > 0 && j > 0 {
		result += matrix[i-1][j-1]
	}
	if i > 0 && j < len(matrix[i])-1 {
		result += matrix[i-1][j+1]
	}
	if i < len(matrix)-1 && j > 0 {
		result += matrix[i+1][j-1]
	}
	if i < len(matrix)-1 && j < len(matrix[i])-1 {
		result += matrix[i+1][j+1]
	}
	return result >= 4
}

func Part2(input string) (string, error) {
	matrix := make([][]int, 0)
	rows := strings.Split(input, "\n")
	result := 0

	for _, row := range rows {
		matrix = append(matrix, make([]int, 0))
		for _, char := range row {
			if char == '@' {
				matrix[len(matrix)-1] = append(matrix[len(matrix)-1], 1)
			} else {
				matrix[len(matrix)-1] = append(matrix[len(matrix)-1], 0)
			}
		}
	}

	queue := make([][2]int, 0)

	for i, row := range matrix {
		for j, cell := range row {
			if cell == 1 {
				if !hasMoreThanFourAdjacentPapers(i, j, matrix) {
					matrix[i][j] = 0
					result++
					queue = append(queue, [2]int{i, j})
				}
			}
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		r, c := curr[0], curr[1]

		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
					if matrix[nr][nc] == 1 {
						if !hasMoreThanFourAdjacentPapers(nr, nc, matrix) {
							matrix[nr][nc] = 0
							result++
							queue = append(queue, [2]int{nr, nc})
						}
					}
				}
			}
		}
	}

	return strconv.Itoa(result), nil
}
