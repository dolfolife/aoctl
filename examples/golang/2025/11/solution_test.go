package main

import "testing"

func TestPart1(t *testing.T) {
	input := `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
	expected := "5"

	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Part1 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	expected := "2"

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Part2 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
