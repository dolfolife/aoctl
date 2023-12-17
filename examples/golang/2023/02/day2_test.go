package day2

import (
	"testing"

	"github.com/dolfolife/aoctl/pkg/puzzle"
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

func (suite *Day2Suite) TestDayTwoPart1TrebuchetCalculateCalibration() {
	suite.SubjectPart1.Puzzle = puzzle.PuzzlePart{
		RawInput: []byte(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`),
	}
	actualValues, errs := suite.SubjectPart1.Solve()
	if errs != nil {
		suite.T().Errorf("Unexpected error: %s\n", errs)
	}

	assert.Equal(suite.T(), "8", actualValues)
}
func (suite *Day2Suite) TestDayTwoPart2TrebuchetCalculateCalibration() {
	suite.SubjectPart2.Puzzle = puzzle.PuzzlePart{
		RawInput: []byte(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`),
	}
	actualValues, errs := suite.SubjectPart2.Solve()
	if errs != nil {
		suite.T().Errorf("Unexpected error: %s\n", errs)
	}

	assert.Equal(suite.T(), "2286", actualValues)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(Day2Suite))
}
