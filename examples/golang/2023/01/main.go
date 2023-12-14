package day1

import (
	"fmt"

	"path/filepath"

	"github.com/dolfolife/aoctl/pkg/aoc"
	"github.com/dolfolife/aoctl/pkg/puzzle"
)

func Solve() {
	aocConfig := aoc.GetAoCConfig()

	file := filepath.Join(aocConfig.ProjectPath, "01/puzzle.yaml")
	inputfile1 := filepath.Join(aocConfig.ProjectPath, "01/input.txt")
	inputfile2 := filepath.Join(aocConfig.ProjectPath, "01/input.txt")
	p := puzzle.NewPuzzleFromCache(file, []string{inputfile1, inputfile2})

	part1 := Day1Part1Solver{}
	part1.Puzzle = p.Puzzles[0]
	anwser1, err := part1.Solve()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(anwser1)

	part2 := Day1Part2Solver{}
	part2.Puzzle = p.Puzzles[1]
	anwser2, err := part2.Solve()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(anwser2)
}
