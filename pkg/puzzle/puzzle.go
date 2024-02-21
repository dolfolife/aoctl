package puzzle

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type PuzzleStatus string

type PuzzleMetadata struct {
	Day   string `yaml:"day"`
	Title string `yaml:"title"`
	Year  string `yaml:"year"`
}

type PuzzlePart struct {
	Answer      string       `yaml:"answer,omitempty"`
	Description string       `yaml:"description"`
	Status      PuzzleStatus `yaml:"status"`

	RawInput []byte
}

type Puzzle struct {
	Metadata PuzzleMetadata `yaml:"metadata"`
	Puzzles  []PuzzlePart   `yaml:"puzzles"`
}

func NewPuzzleFromHTML(day string, year string, htmlString string, input []byte) Puzzle {
	return Puzzle{
		Metadata: PuzzleMetadata{
			Day:   day,
			Year:  year,
			Title: getTitleFromBody(htmlString),
		},
		Puzzles: getPuzzlePartsFromHTMLString(htmlString),
	}
}

// The field status is a collection of status and we need to validate that the
// status is in the set of valid statuses
func (p *Puzzle) ParseFields() error {
	mapStatus := map[string]PuzzleStatus{
		"UNSOLVED":    Unsolved,
		"SOLVED":      Solved,
		"UNREACHABLE": Unreachable,
	}
	for i, puzzle := range p.Puzzles {
		status := strings.ToUpper(string(puzzle.Status))
		if _, ok := mapStatus[status]; ok {
			p.Puzzles[i].Status = mapStatus[status]
		} else {
			errorMessage := fmt.Sprintf("cannot parse Puzzle Part %d", i)
			return NewError(ErrInvalidStatus, errors.New(errorMessage))
		}
	}
	return nil
}

func NewPuzzleFromCache(filepath string, inputFilepath []string) (Puzzle, error) {
	var puzzle Puzzle
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Error trying to read the YAML file err =  #%v ", err)
		return Puzzle{}, err
	}
	err = yaml.Unmarshal(yamlFile, &puzzle)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return Puzzle{}, err
	}

	// yaml.Umarshall does not have validation on sets like the status field
	// We need to map the status of the puzzle
	err = puzzle.ParseFields()
	if err != nil {
		log.Fatalf("Error trying to parse the Puzzle err =  #%v ", err)
		return Puzzle{}, err
	}

	for i, inputFile := range inputFilepath {
		rawInput, err := os.ReadFile(inputFile)
		if err != nil {
			log.Fatalf("Error trying to read the input for Puzzle Part %d err   #%v ", i, err)
			return Puzzle{}, err
		}
		// we need to delete the last byte of the input because it is a newline or EOF
		rawInput = rawInput[:len(rawInput)-1]
		puzzle.Puzzles[i].RawInput = rawInput
	}
	return puzzle, nil
}

func getTitleFromBody(body string) string {
	parts := strings.Split(body, "---")
	titleParts := strings.Split(parts[1], ":")
	return strings.TrimSpace(titleParts[1])
}

func getPuzzlePartsFromHTMLString(body string) []PuzzlePart {
	puzzleParts := make([]PuzzlePart, 0)

	parts := strings.Split(body, "---")

	if len(parts) < 3 {
		log.Fatal("Description of the puzzle is not found")
		return puzzleParts
	}
	puzzleParts = append(puzzleParts, PuzzlePart{
		Answer:      "",
		Status:      "UNSOLVED",
		Description: parts[2],
	})

	return puzzleParts
}

type PuzzleSolver[T any] struct {
	Puzzle         Puzzle
	NormalizeInput func(string) T
	Solve          func() Response
}

type Response struct {
	Value string
	Error error
}
