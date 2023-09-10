package day_1

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDayOneNoTimeForaTaxicab1(t *testing.T) {
    expectedValues := []string{"5", ""}
    actualValues := Solve("R2, L3")
    assert.Equal(t, expectedValues[0], actualValues.Part1)
}

func TestDayOneNoTimeForaTaxicab2(t *testing.T) {
    expectedValues := []string{"2", ""}
    actualValues := Solve("R2, R2, R2")
    assert.Equal(t, expectedValues[0], actualValues.Part1)
}

func TestDayOneNoTimeForaTaxicab3(t *testing.T) {
    expectedValues := []string{"12", ""}
    actualValues := Solve("R5, L5, R5, R3")
    assert.Equal(t, expectedValues[0], actualValues.Part1)
}

func TestDayOneNoTimeForaTaxicab4(t *testing.T) {
    expectedValues := []string{"2", ""}
    actualValues := Solve("L2, L2, L2")
    assert.Equal(t, expectedValues[0], actualValues.Part1)
}

func TestDayOneNoTimeForaTaxicabVisitedTwice(t *testing.T) {
    expectedValues := []string{"8", "4"}
    actualValues := Solve("R8, R4, R4, R8")
    assert.Equal(t, expectedValues[0], actualValues.Part1)
    assert.Equal(t, expectedValues[1], actualValues.Part2)
}
