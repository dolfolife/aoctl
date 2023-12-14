package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
func part1(input []string) (string, error) {
	resultArray := []int{}
	for _, line := range input {
		firstDigit, lastDigit := '0', '0'
		for _, char := range line {
			if isDigit(char) {
				firstDigit = char
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if isDigit(rune(line[i])) {
				lastDigit = rune(line[i])
				break
			}
		}
		newInt, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		if err != nil {
			return "", err
		}

		resultArray = append(resultArray, newInt)
	}
	sum := 0
	for _, value := range resultArray {
		sum += value
	}
	return fmt.Sprint(sum), nil
}

type Day1Part1Solver struct {
	Puzzle puzzle.PuzzlePart
}

func (s *Day1Part1Solver) NormalizeInput(input string) []string {
	return strings.Split(input, "\n")
}

func (s *Day1Part1Solver) Solve() (string, error) {
	return part1(s.NormalizeInput(string(s.Puzzle.RawInput)))
}
