package main

import "testing"

func TestPart1(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	expected := "357"

	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	expected := "3121910778619"

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
