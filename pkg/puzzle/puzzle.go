package puzzle

import (
    "fmt"
)

type Puzzle[T any] struct {
    part1 func(T) (string, error)
    part2 func(T) (string, error)
    normalizeInput func(string) T
}

func (p Puzzle[T]) SolvePart(part int,input string) string {
    var result string
    var err error

    if part == 1 {
        result, err = p.part1(p.normalizeInput(input))
    } else {
        result, err = p.part2(p.normalizeInput(input))
    }

    if err != nil {
        fmt.Printf("there was an error on part %d: %s \n", part, err)
        fmt.Printf("...skipping part %d...\n", part)
        return ""
    }

    return result
}

func (p Puzzle[T]) NormalizeInput(input string) T {
    return p.normalizeInput(input)
}

func NewPuzzleSolver[T any](name string, normalizeInput func(string) T, part1 func(T) (string, error), part2 func(T) (string, error)) Puzzle[T] {
    return Puzzle[T] {
        part1: part1,
        part2: part2,
        normalizeInput: normalizeInput,
    }
}

type PuzzleSolution[T any] struct {
    Part1 string
    Part2 string

    Error error
}
