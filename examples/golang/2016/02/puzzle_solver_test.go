package day2

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type Day2Suite struct {
    suite.Suite
    SubjectPart1 Day2Part1Solver
    SubjectPart2 Day2Part2Solver
}

func (suite *Day2Suite) SetupTest() {
    suite.SubjectPart1 = Day2Part1Solver{}
    suite.SubjectPart2 = Day2Part2Solver{}
}

func (suite *Day2Suite) TestBathroomSecurityPart1() {
    input := `ULL
RRDDD
LURDL
UUUUD`

    suite.SubjectPart1.Puzzle.RawInput = []byte(input)
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }

    assert.Equal(suite.T(), "1985", actualValues)

}

func (suite *Day2Suite) TestBathroomSecurityPart2() {
    input := `ULL
RRDDD
LURDL
UUUUD`
    suite.SubjectPart2.Puzzle.RawInput = []byte(input)
    actualValues, errs := suite.SubjectPart2.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "5DB3", actualValues)
}

func TestSuite(t *testing.T) {
    suite.Run(t, new(Day2Suite))
}

