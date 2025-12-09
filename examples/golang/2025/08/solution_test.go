package main

import "testing"

func TestPart1(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	expected := "40"

	Part1Limit = 10
	defer func() { Part1Limit = 1000 }()

	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}

	if result != expected {
		t.Errorf("Part1 result mismatch: expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	expected := "25272"

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}

	if result != expected {
		t.Errorf("Part2 result mismatch: expected %s, got %s", expected, result)
	}
}
