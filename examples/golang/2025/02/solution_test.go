package main

import "testing"

func TestPart1(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`
	expected := "1227775554"

	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Part1 = %q, want %q", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
		1698522-1698528,446443-446449,38593856-38593862,565653-565659,
		824824821-824824827,2121212118-2121212124`
	expected := "4174379265"

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Part2 = %q, want %q", result, expected)
	}
}
