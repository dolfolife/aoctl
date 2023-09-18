package _day02

import (
    "strings"
    "fmt"

    "github.com/dolfolife/aoctl/pkg/puzzle"
)

func solvePart1(commands []string) (string, error) {
    position := 5
    var code []string
    for _, command := range commands {
        for _, c := range command {
            switch c {
            case 'U':
                if position > 3 {
                    position -= 3
                }
            case 'D':
                if position < 7 {
                    position += 3
                }
            case 'L':
                if position%3 != 1 {
                    position -= 1
                }
            case 'R':
                if position%3 != 0 {
                    position += 1
                }
            }
        }
        code = append(code, fmt.Sprint(position))
    }

    return strings.Join(code, ""), nil
}

func solvePart2(commands []string) (string, error) {
    position := 5
    var code []string
    for _, command := range commands {
        for _, c := range command {
            switch c {
            case 'U':
                if position != 1 && position != 2 && position != 4 && position != 5 && position != 9 {
                    if position == 3 || position == 13 {
                        position -= 2
                    } else {
                        position -= 4
                    }
                }
            case 'D':
                if position != 5 && position != 9 && position != 10 && position != 12 && position != 13 {
                    if position == 1 || position == 11 {
                        position += 2
                    } else {
                        position += 4
                    }
                }
            case 'L':
                if position != 1 && position != 2 && position != 5 && position != 10 && position != 13 {
                    position -= 1
                }
            case 'R':
                if position != 1 && position != 4 && position != 9 && position != 12 && position != 13 {
                    position += 1
                }
            }
        }
        code = append(code, fmt.Sprintf("%X", position))
    }

    return strings.Join(code, ""), nil
}

func normalizeInput(input string) []string {
    return strings.Split(input, "\n")
}

func GetPuzzleSolver() puzzle.Puzzle[[]string] {
    return puzzle.NewPuzzleSolver("day02", normalizeInput, solvePart1, solvePart2)
}

func Solve(fileInputContent string) puzzle.PuzzleSolution[[]string] {

    puzzleSolver := GetPuzzleSolver()
    return puzzle.PuzzleSolution[[]string]{
        Part1: puzzleSolver.SolvePart(1, fileInputContent),
        Part2: puzzleSolver.SolvePart(2, fileInputContent),
    }
}
