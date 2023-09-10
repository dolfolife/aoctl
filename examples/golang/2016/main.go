package main

import (
    "os"
    "fmt"

    day_1 "github.com/dolfolife/adventofcode/examples/golang/2016/01"
)

func main() {
    input, err := os.ReadFile("./01/input.txt")
    if err != nil {
        os.Exit(1)
    }

    // We need to delete the last rune of the input since it is a newline
    answers := day_1.Solve(string(input[:len(input)-1]))

    fmt.Printf("Part 1: %s\n", answers.Part1)
    fmt.Printf("Part 2: %s\n", answers.Part2)
}
