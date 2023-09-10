package _day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dolfolife/adventofcode/pkg/math"
	"github.com/dolfolife/adventofcode/pkg/puzzle"
)

func solvePart1(input [][]int) (string, error) {
    validTriangles := 0

    for _, line := range input {
        if math.IsValidTriangle(line[0], line[1], line[2]) {
            validTriangles += 1
        }
    }
    return fmt.Sprint(validTriangles), nil
}

func solvePart2(input [][]int) (string, error) {
    validTriangles := 0
    for i := 0; i < len(input); i += 3 {
        for j := 0; j < 3; j++ {
            if math.IsValidTriangle(input[i][j], input[i+1][j], input[i+2][j]) {
                validTriangles += 1
            }
        }
    }

    return fmt.Sprint(validTriangles), nil
}

func normalizeInput(input string) [][]int {
    result := make([][]int, 0)
    lines := strings.Split(input, "\n") 
    for _, triangleSizes := range lines {
        rawValues := strings.Split(triangleSizes, " ")
        line := make([]int, 0)
        for _, value := range rawValues {
            v, _ := strconv.Atoi(value)
            if v > 0 {
                line = append(line, v)
            }
        }
        result = append(result, line)
    }
    return result
}

func GetPuzzleSolver() puzzle.Puzzle[[][]int] {
    return puzzle.NewPuzzleSolver("day03", normalizeInput, solvePart1, solvePart2)
}

func Solve(input string) puzzle.PuzzleSolution[[][]int] {
    puzzleSolver := GetPuzzleSolver()

    return puzzle.PuzzleSolution[[][]int]{
        Part1: puzzleSolver.SolvePart(1, input),
        Part2: puzzleSolver.SolvePart(2, input),
    }
}
