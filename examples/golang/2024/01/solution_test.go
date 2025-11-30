package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := "11"
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
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := "31"
	result, err := Part2(input)
	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
