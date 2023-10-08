package day1

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/dolfolife/aoctl/pkg/puzzle"
)

type Day1Suite struct {
    suite.Suite
    SubjectPart1 Day1Part1Solver
    SubjectPart2 Day1Part2Solver
}

func (suite *Day1Suite) SetupTest() {
    suite.SubjectPart1 = Day1Part1Solver{}
    suite.SubjectPart2 = Day1Part2Solver{}
}

func (suite *Day1Suite) TestDayOnePart1NoTimeForaTaxicab1() {
    suite.SubjectPart1.Puzzle = puzzle.PuzzlePart{
        RawInput: []byte("R2, L3"),
    }
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }

    assert.Equal(suite.T(), "5", actualValues)
}

func (suite *Day1Suite) TestDayOnePart1NoTimeForaTaxicab2() {
    suite.SubjectPart1.Puzzle = puzzle.PuzzlePart{
        RawInput: []byte("R2, R2, R2"),
    }
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "2", actualValues)
}

func (suite *Day1Suite) TestDayOnePart1NoTimeForaTaxicab3() {
    suite.SubjectPart1.Puzzle = puzzle.PuzzlePart{
        RawInput: []byte("R5, L5, R5, R3"),
    }
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "12", actualValues)
}

func (suite *Day1Suite) TestDayOnePart1NoTimeForaTaxicab4() {
    suite.SubjectPart1.Puzzle = puzzle.PuzzlePart{
        RawInput: []byte("L2, L2, L2"),
    }
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "2", actualValues)
}

func (suite *Day1Suite) TestDayOnePart1NoTimeForaTaxicabVisitedTwice() {
    suite.SubjectPart1.Puzzle = puzzle.PuzzlePart{
        RawInput: []byte("R8, R4, R4, R8"),
    }
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "8", actualValues)
}

func (suite *Day1Suite) TestDayOnePart2NoTimeForaTaxicabVisitedTwice() {
    suite.SubjectPart2.Puzzle = puzzle.PuzzlePart{
        RawInput: []byte("R8, R4, R4, R8"),
    }
    actualValues, errs := suite.SubjectPart2.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "4", actualValues)
}

func TestSuite(t *testing.T) {
    suite.Run(t, new(Day1Suite))
}

