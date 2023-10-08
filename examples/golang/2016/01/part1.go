package day1

import (
    "strings"
    "strconv"

    . "github.com/dolfolife/aoctl/pkg/math"
    "github.com/dolfolife/aoctl/pkg/puzzle"
)

func part1(input []string) (string, error) {
    var orientation = []rune{ 'N', 'E', 'S', 'W' }
    posCurrentOrientation := 0

    var pos = Point{X: 0, Y: 0}
    for _, v := range input {
        if v[0] == 'R' {
            posCurrentOrientation = (posCurrentOrientation + 1) % len(orientation)
        } else {
            posCurrentOrientation = (posCurrentOrientation - 1) % len(orientation)
        }

        if posCurrentOrientation < 0 {
            posCurrentOrientation = len(orientation) + posCurrentOrientation
        }
        steps, _ := strconv.Atoi(v[1:])
        segment := newSegmentPath(orientation[posCurrentOrientation], pos, float64(steps))
        pos = segment.B
    }
    return CalDistancePoints(Point{X: 0, Y: 0}, pos), nil
}

type Day1Part1Solver struct {
    Puzzle puzzle.PuzzlePart
}

func (s *Day1Part1Solver) NormalizeInput(input string) []string {
    return strings.Split(input, ", ")
}

func (s *Day1Part1Solver) Solve() (string, error) {
    return part1(s.NormalizeInput(string(s.Puzzle.RawInput)))
}

