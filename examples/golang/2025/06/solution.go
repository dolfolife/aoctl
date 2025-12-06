package main

import (
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	linesStr := strings.Split(input, "\n")
	lines := make([][]string, 0)
	for _, line := range linesStr {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			lines = append(lines, fields)
		}
	}
	result := 0
	for i := 0; i < len(lines[0]); i++ {
		mathNumbers := make([]int, 0)
		for j := 0; j < len(lines)-1; j++ {
			mathNumber, _ := strconv.Atoi(lines[j][i])
			mathNumbers = append(mathNumbers, mathNumber)
		}
		mathOperation := lines[len(lines)-1][i]
		switch mathOperation {
		case "+":
			localResult := 0
			for _, mathNumber := range mathNumbers {
				localResult += mathNumber
			}
			result += localResult
		case "*":
			localResult := 1
			for _, mathNumber := range mathNumbers {
				localResult *= mathNumber
			}
			result += localResult
		}
	}
	return strconv.Itoa(result), nil
}

func Part2(input string) (string, error) {
	linesStr := strings.Split(input, "\n")
	lines := make([][]string, 0)
	for _, line := range linesStr {
		lines = append(lines, strings.Split(line, ""))
	}
	result := 0
	maxLength := len(lines[0])
	mathNumbers := make([]int, 0)

	for i := maxLength - 1; i >= 0; i-- {
		mathNumberString := ""
		for j := 0; j < len(lines)-1; j++ {
			mathNumberString += lines[j][i]
		}
		mathNumberString = strings.TrimSpace(mathNumberString)
		if mathNumberString == "" {
			result += calculateResult(mathNumbers, lines[len(lines)-1][i+1])
			mathNumbers = make([]int, 0)
		} else {
			mathNumber, _ := strconv.Atoi(mathNumberString)
			mathNumbers = append(mathNumbers, mathNumber)
		}
	}
	result += calculateResult(mathNumbers, lines[len(lines)-1][0])
	return strconv.Itoa(result), nil
}

func calculateResult(mathNumbers []int, mathOperation string) int {
	switch mathOperation {
	case "+":
		localResult := 0
		for _, mathNumber := range mathNumbers {
			localResult += mathNumber
		}
		return localResult
	case "*":
		localResult := 1
		for _, mathNumber := range mathNumbers {
			localResult *= mathNumber
		}
		return localResult
	}
	return 0
}
