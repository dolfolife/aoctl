package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

func startWithDigit(input string) (string, bool) {
	digitsWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, word := range digitsWords {
		if strings.HasPrefix(input, word) {
			return word, true
		}
	}
	return "", false
}

func part2(input []string) (string, error) {
	resultArray := []int{}
	for _, line := range input {
		//fix line
		digitsWords := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
		firstDigit, lastDigit := '0', '0'
		for i, char := range line {
			if isDigit(char) {
				firstDigit = char
				break
			}
			digitVal, isDigitString := startWithDigit(line[i:])
			if isDigitString {
				firstDigit = rune(digitsWords[digitVal] + '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if isDigit(rune(line[i])) {
				lastDigit = rune(line[i])
				break
			}
			digitVal, isDigitString := startWithDigit(line[i:])
			if isDigitString {
				lastDigit = rune(digitsWords[digitVal] + '0')
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

type Day1Part2Solver struct {
	Puzzle puzzle.PuzzlePart
}

func (s *Day1Part2Solver) NormalizeInput(input string) []string {
	return strings.Split(input, "\n")
}

func (s *Day1Part2Solver) Solve() (string, error) {
	return part2(s.NormalizeInput(string(s.Puzzle.RawInput)))
}
