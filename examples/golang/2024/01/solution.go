package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	metricsOne := []int{}
	metricsTwo := []int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		digitsInLine := strings.Split(line, " ")
		if len(digitsInLine) != 4 {
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

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	metricsOne := []int{}
	metricsTwo := []int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		digitsInLine := strings.Split(line, " ")
		if len(digitsInLine) != 4 {
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
