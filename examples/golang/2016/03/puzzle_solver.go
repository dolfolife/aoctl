package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dolfolife/aoctl/pkg/math"
	"github.com/dolfolife/aoctl/pkg/puzzle"
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

type Day3Part1Solver struct {
    Puzzle puzzle.PuzzlePart
}

func (s *Day3Part1Solver) NormalizeInput(input string) [][]int {
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

func (s *Day3Part1Solver) Solve() (string, error) {
    return solvePart1(s.NormalizeInput(string(s.Puzzle.RawInput)))
}

type Day3Part2Solver struct {
    Puzzle puzzle.PuzzlePart
}

func (s *Day3Part2Solver) NormalizeInput(input string) [][]int {
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

func (s *Day3Part2Solver) Solve() (string, error) {
    return solvePart2(s.NormalizeInput(string(s.Puzzle.RawInput)))
}
