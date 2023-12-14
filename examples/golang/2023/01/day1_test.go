package day1

import (
	"testing"

	"github.com/dolfolife/aoctl/pkg/puzzle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

func (suite *Day1Suite) TestDayOnePart1TrebuchetCalculateCalibration() {
	suite.SubjectPart1.Puzzle = puzzle.PuzzlePart{
		RawInput: []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`),
	}
	actualValues, errs := suite.SubjectPart1.Solve()
	if errs != nil {
		suite.T().Errorf("Unexpected error: %s\n", errs)
	}

	assert.Equal(suite.T(), "142", actualValues)
}
func (suite *Day1Suite) TestDayOnePart2TrebuchetCalculateCalibration() {
	suite.SubjectPart2.Puzzle = puzzle.PuzzlePart{
		RawInput: []byte(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`),
	}
	actualValues, errs := suite.SubjectPart2.Solve()
	if errs != nil {
		suite.T().Errorf("Unexpected error: %s\n", errs)
	}

	assert.Equal(suite.T(), "281", actualValues)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(Day1Suite))
}
