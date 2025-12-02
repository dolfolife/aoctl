package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	ranges := strings.Split(input, ",")
	result := 0
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		minStr, maxStr := parts[0], parts[1]
		minInt, err := strconv.Atoi(minStr)
		if err != nil {
			return "", err
		}
		maxInt, err := strconv.Atoi(maxStr)
		if err != nil {
			return "", err
		}
		for i := minInt; i <= maxInt; i++ {
			if isDoubleString(i) {
				result += i
			}
		}
	}
	return fmt.Sprint(result), nil
}

func isDoubleString(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	return s[:len(s)/2] == s[len(s)/2:]
}

func Part2(input string) (string, error) {
	ranges := strings.Split(input, ",")
	result := 0
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		minStr, maxStr := parts[0], parts[1]
		minInt, err := strconv.Atoi(minStr)
		if err != nil {
			return "", err
		}
		maxInt, err := strconv.Atoi(maxStr)
		if err != nil {
			return "", err
		}
		for i := minInt; i <= maxInt; i++ {
			if isFromRepetitions(i) {
				result += i
			}
		}
	}
	return fmt.Sprint(result), nil
}

func isFromRepetitions(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)
	for patLen := 1; patLen <= length/2; patLen++ {
		if length%patLen != 0 {
			continue
		}
		pattern := s[:patLen]
		isPattern := true
		for i := patLen; i < length; i += patLen {
			if s[i:i+patLen] != pattern {
				isPattern = false
				break
			}
		}
		if isPattern {
			return true
		}
	}
	return false
}
