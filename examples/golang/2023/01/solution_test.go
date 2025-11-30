package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	expected := "142"
	result, err := Part1(input)
	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	expected := "281"
	result, err := Part2(input)
	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
