package day1

import (
    "strings"
    "strconv"
    "errors"

    . "github.com/dolfolife/aoctl/pkg/math"
    "github.com/dolfolife/aoctl/pkg/puzzle"
)

func part2(input []string) (string, error) {
    var orientation = []rune{ 'N', 'E', 'S', 'W' }
    posCurrentOrientation := 0

   path := make([]Segment, 0)

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
        newPath := newSegmentPath(orientation[posCurrentOrientation], pos, float64(steps))

        for _, segment := range path {
            intersect, intersectionPoint := CalIntersectionPoint(segment.A, segment.B, newPath.A, newPath.B)
            if intersect {
                return CalDistancePoints(Point{X: 0, Y: 0}, intersectionPoint), nil
            }
        }
        path = append(path, newPath) 
        pos = newPath.B
    }
    return "", errors.New("No intersection found")
}

type Day1Part2Solver struct {
    Puzzle puzzle.PuzzlePart
}

func (s *Day1Part2Solver) NormalizeInput(input string) []string {
    return strings.Split(input, ", ")
}

func (s *Day1Part2Solver) Solve() (string, error) {
    return part2(s.NormalizeInput(string(s.Puzzle.RawInput)))
}

