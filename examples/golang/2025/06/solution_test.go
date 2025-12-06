package main

import "testing"

func TestPart1(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`
	expected := "4277556"
	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Part1 expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	expected := "3263827"
	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Part2 expected %s, got %s", expected, result)
	}
}
