package main

import "testing"

func TestPart1(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	expected := "3"

	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}

	if result != expected {
		t.Errorf("Part1 = %q, want %q", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	expected := "14"

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}

	if result != expected {
		t.Errorf("Part2 = %q, want %q", result, expected)
	}
}

func TestPart2_Encapsulation(t *testing.T) {
	input := `10-20
5-25

1
`
	expected := "21"

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}

	if result != expected {
		t.Errorf("Part2 = %q, want %q", result, expected)
	}
}
