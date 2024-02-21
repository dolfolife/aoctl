package puzzle

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePuzzleFromCache(t *testing.T) {
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

func TestParsePuzzleFromHTML(t *testing.T) {

	ROOT_DIR := os.Getenv("PWD")
	testFilePath := filepath.Join(ROOT_DIR, "test_data/puzzle_with_one_question.txt")
	file, err := os.ReadFile(testFilePath)

	if err != nil {
		t.Fail()
	}

	puzzle := NewPuzzleFromHTML("day", "year", string(file), []byte{})

	assert.Equal(t, "day", puzzle.Metadata.Day)
	assert.Equal(t, "year", puzzle.Metadata.Year)
	assert.Equal(t, "Trebuchet?!", puzzle.Metadata.Title)
	assert.Equal(t, 1, len(puzzle.Puzzles))
	assert.Equal(t, 1827, len(puzzle.Puzzles[0].Description))

}
