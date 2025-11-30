package day1

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

type Day1Part1Solver struct {
	Puzzle puzzle.PuzzlePart
}

func part1(input []string) (string, error) {
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
	for index, value := range metricsOne {
		localDiff := int(math.Abs(float64(value - metricsTwo[index])))
		result += localDiff
	}
	return fmt.Sprint(result), nil
}

func (s *Day1Part1Solver) NormalizeInput(input string) []string {
	return strings.Split(input, "\n")
}

func (s *Day1Part1Solver) Solve() (string, error) {
	return part1(s.NormalizeInput(string(s.Puzzle.RawInput)))
}
