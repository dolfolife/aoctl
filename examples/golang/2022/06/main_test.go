package main

import "testing"

func TestSolution1(t *testing.T) {
	testInput := map[string]int{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}
	for input, expectedOutput := range testInput {
		actual := FindFirstMarker(4, []byte(input))
		if actual != expectedOutput {
			t.Errorf("wanted %d; but got %d", actual, expectedOutput)
		}
	}
}

func TestSolution2(t *testing.T) {
	testInput := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}
	for input, expectedOutput := range testInput {
		actual := FindFirstMarker(14, []byte(input))
		if actual != expectedOutput {
			t.Errorf("wanted %d; but got %d", actual, expectedOutput)
		}
	}
}
