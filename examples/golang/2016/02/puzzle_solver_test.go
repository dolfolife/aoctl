package _day02

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBathroomSecurityPart1(t *testing.T) {
    subject := GetPuzzleSolver()
    input := `ULL
RRDDD
LURDL
UUUUD`
    assert.Equal(t, "1985", subject.SolvePart(1, input))
}

func TestBathroomSecurityPart2(t *testing.T) {
    subject := GetPuzzleSolver()
    input := `ULL
RRDDD
LURDL
UUUUD`
    assert.Equal(t, "5DB3", subject.SolvePart(2, input))
}
