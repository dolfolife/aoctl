package _day03

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestSquaresWithThreeSidesPart1(t *testing.T) {

    subject := GetPuzzleSolver()

    input := `5 10 25`
    assert.Equal(t, "0", subject.SolvePart(1, input))

    input = `5 10 25
1 1 1
5 10 25`
    assert.Equal(t, "1", subject.SolvePart(1, input))
}

func TestSquaresWithThreeSidesPart2(t *testing.T) {

    subject := GetPuzzleSolver()

    input := `1 10 25
1 1 1
1 10 25`
    assert.Equal(t, "3", subject.SolvePart(2, input))
    
    input = `1 10 25
1 1 1
1 10 25
1 1 1
1 1 1
1 1 1
1 1 1
10 100 10
1 1 1`
    assert.Equal(t, "6", subject.SolvePart(2, input))
}
