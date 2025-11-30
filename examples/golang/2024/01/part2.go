package day1

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

type Day1Part2Solver struct {
	Puzzle puzzle.PuzzlePart
}

func part2(input []string) (string, error) {
	metricsOne := []int{}
	metricsTwo := []int{}
	for _, line := range input {
		digitsInLine := strings.Split(line, " ")
		if len(digitsInLine) != 4 {
			fmt.Println(len(digitsInLine))
			return "", errors.New("not enough digits in line")
		}
		firstDigit, err := strconv.Atoi(digitsInLine[0])
		if err != nil {
			return "", err
		}
		metricsOne = append(metricsOne, firstDigit)
		secondDigit, err := strconv.Atoi(digitsInLine[3])
		if err != nil {
			return "", err
		}
		metricsTwo = append(metricsTwo, secondDigit)
	}
	sort.Ints(metricsOne)
	sort.Ints(metricsTwo)
	result := 0
	secondIndex := 0
	var prevValue int = -1
	var prevScore int = 0

	for i, value := range metricsOne {
		if i > 0 && value == prevValue {
			result += prevScore
			continue
		}

		for secondIndex < len(metricsTwo) && metricsTwo[secondIndex] < value {
			secondIndex++
		}

		count := 0
		for secondIndex < len(metricsTwo) && metricsTwo[secondIndex] == value {
			count++
			secondIndex++
		}

		score := value * count
		result += score

		prevValue = value
		prevScore = score
	}
	return fmt.Sprint(result), nil
}

func (s *Day1Part2Solver) NormalizeInput(input string) []string {
	return strings.Split(input, "\n")
}

func (s *Day1Part2Solver) Solve() (string, error) {
	return part2(s.NormalizeInput(string(s.Puzzle.RawInput)))
}
