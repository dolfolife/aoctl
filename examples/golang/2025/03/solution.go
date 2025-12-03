package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findMaxJoltage(parts []string) (string, int) {
	maxJoltage := 0
	maxJoltageString := ""
	maxJoltageIndex := 0
	for i, part := range parts {
		joltage, _ := strconv.Atoi(part)
		if joltage > maxJoltage {
			maxJoltage = joltage
			maxJoltageString = part
			maxJoltageIndex = i
		}
	}
	return maxJoltageString, maxJoltageIndex
}
func Part1(input string) (string, error) {
	result := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "")

		maxJoltageString, maxJoltageIndex := findMaxJoltage(parts[:len(parts)-1])
		secondMaxJoltageString, _ := findMaxJoltage(parts[maxJoltageIndex+1:])

		joltage, _ := strconv.Atoi(maxJoltageString + secondMaxJoltageString)
		result += joltage
	}
	return fmt.Sprint(result), nil
}

func Part2(input string) (string, error) {
	total := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		stack := make([]int, 0, 12)
		digits := make([]int, 0, len(line))
		for _, r := range line {
			if r >= '0' && r <= '9' {
				digits = append(digits, int(r-'0'))
			}
		}

		n := len(digits)
		k := 12
		if n < k {
			continue
		}

		for i, digit := range digits {
			for len(stack) > 0 && stack[len(stack)-1] < digit && len(stack)-1+(n-i) >= k {
				stack = stack[:len(stack)-1]
			}
			if len(stack) < k {
				stack = append(stack, digit)
			}
		}

		val := 0
		for _, d := range stack {
			val = val*10 + d
		}
		total += val
	}
	return strconv.Itoa(total), nil
}
