package day_1

import (
    "strings"
    "strconv"
    "errors"

    "github.com/dolfolife/adventofcode/pkg/puzzle"
    . "github.com/dolfolife/adventofcode/pkg/math"
)


func newSegmentPath(orientation rune, p Point, steps float64) Segment {
    startingPoint := p
    switch orientation {
    case 'N':
        p.Y += steps
        startingPoint.Y += 1
    case 'E':
        p.X += steps
        startingPoint.X += 1
    case 'S':
        p.Y -= steps
        startingPoint.Y -= 1
    case 'W':
        p.X -= steps
        startingPoint.X -= 1
    }
    return Segment{A: startingPoint, B: p}
}

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

func normalize(input string) []string {
    return strings.Split(input, ", ")
}

func Solve(fileInputContent string) puzzle.PuzzleSolution[[]string] {
    puzzleSolver := puzzle.NewPuzzleSolver("day_1", normalize, part1, part2)
    return puzzle.PuzzleSolution[[]string] {
        Part1: puzzleSolver.SolvePart(1, fileInputContent),
        Part2: puzzleSolver.SolvePart(2, fileInputContent),
    }
}
