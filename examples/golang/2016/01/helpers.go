package day1

import (

    . "github.com/dolfolife/aoctl/pkg/math"
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
