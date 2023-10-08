package day3

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

func (suite *Day3Suite) TestSquaresWithThreeSidesPart1_1() {

    input := `5 10 25`
    suite.SubjectPart1.Puzzle.RawInput = []byte(input)
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "0", actualValues)
}

func (suite *Day3Suite) TestSquaresWithThreeSidesPart1_2() {
    input := `5 10 25
1 1 1
5 10 25`
    suite.SubjectPart1.Puzzle.RawInput = []byte(input)
    actualValues, errs := suite.SubjectPart1.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "1", actualValues)
}

func (suite *Day3Suite) TestSquaresWithThreeSidesPart2_1() {
    input := `1 10 25
1 1 1
1 10 25`
    suite.SubjectPart2.Puzzle.RawInput = []byte(input)
    actualValues, errs := suite.SubjectPart2.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "3", actualValues)
}

func (suite *Day3Suite) TestSquaresWithThreeSidesPart2_2() {
    input := `1 10 25
1 1 1
1 10 25
1 1 1
1 1 1
1 1 1
1 1 1
10 100 10
1 1 1`
    suite.SubjectPart2.Puzzle.RawInput = []byte(input)
    actualValues, errs := suite.SubjectPart2.Solve()
    if errs != nil {
        suite.T().Errorf("Unexpected error: %s\n", errs)
    }
    assert.Equal(suite.T(), "6", actualValues)
}

type Day3Suite struct {
    suite.Suite
    SubjectPart1 Day3Part1Solver
    SubjectPart2 Day3Part2Solver
}

func (suite *Day3Suite) SetupTest() {
    suite.SubjectPart1 = Day3Part1Solver{}
    suite.SubjectPart2 = Day3Part2Solver{}
}

func TestSuite(t *testing.T) {
    suite.Run(t, new(Day3Suite))
}

