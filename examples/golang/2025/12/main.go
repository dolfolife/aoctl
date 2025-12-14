package main

import (
	"fmt"
	"log"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

func main() {
	input, err := puzzle.ReadInput("input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	p1, err := Part1(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: %s\n", p1)
}
