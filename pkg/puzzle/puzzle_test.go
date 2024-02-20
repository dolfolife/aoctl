package puzzle

import (
    "testing"
    "os"
    "path/filepath"

    "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    ROOT_DIR := os.Getenv("PWD")
    basicYaml := filepath.Join(ROOT_DIR, "test_data/basic.yaml")
    subject, err := NewPuzzleFromCache(basicYaml, []string{})

    // Basic Yaml has no errors
    assert.Nil(t, err)

    // Metada is set correctly
    assert.Equal(t, subject.Metadata.Year, "year_basic")
    assert.Equal(t, subject.Metadata.Day, "day_basic")
    assert.Equal(t, subject.Metadata.Title, "title_basic")

    // Puzzles check
    assert.Equal(t, len(subject.Puzzles), 2)
    assert.Equal(t, subject.Puzzles[0].Description, "description_1_basic") 
    assert.Equal(t, subject.Puzzles[0].Answer, "answer_1_basic") 
    assert.Equal(t, subject.Puzzles[0].Status, Solved) 
}

