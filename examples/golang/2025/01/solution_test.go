package main

import "testing"

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	expected := "3"
	result, err := Part1(input)
	if err != nil {
		t.Errorf("Part1() error = %v", err)
		return
	}
	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart1_LargeRotations(t *testing.T) {
	input := "R350"
	expected := "1"
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
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	expected := "6"
	result, err := Part2(input)
	if err != nil {
		t.Errorf("Part2() error = %v", err)
		return
	}
	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
