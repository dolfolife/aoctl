package day2

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/dolfolife/aoctl/pkg/puzzle"
    "github.com/dolfolife/aoctl/pkg/aoc"
)

func Solve() {

    aocConfig := aoc.GetAoCConfig()

    file := filepath.Join(aocConfig.ProjectPath, "02/puzzle.yaml")
    inputfile1 := filepath.Join(aocConfig.ProjectPath, "02/input.txt")
    inputfile2 := filepath.Join(aocConfig.ProjectPath, "02/input.txt")
    p, err := puzzle.NewPuzzleFromCache(file, []string{inputfile1, inputfile2})
    
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    part1 := Day2Part1Solver{}
    part1.Puzzle = p.Puzzles[0]
    answer1, err := part1.Solve()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("For the Year %s on the day %s, -- %s\n", p.Metadata.Year, p.Metadata.Day, p.Metadata.Title)
    fmt.Printf("Part 1 -- \n Actual Answer: %s and last recorded answer is %s\n", answer1, p.Puzzles[0].Answer)

    part2 := Day2Part2Solver{}
    part2.Puzzle = p.Puzzles[1]
    answer2, err2 := part2.Solve()
    if err2 != nil {
        fmt.Println(err2)
    }
    fmt.Printf("Part 2 -- \n Actual Answer: %s and last recorded answer is %s\n", answer2, p.Puzzles[1].Answer)
}
