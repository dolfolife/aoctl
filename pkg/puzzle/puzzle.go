package puzzle

import(
    "gopkg.in/yaml.v3"
    "log"
    "os"
)

type PuzzleStatus string

const (
    Unsolved    PuzzleStatus = "UNSOLVED"
    Solved      PuzzleStatus = "SOLVED"
    Unreachable PuzzleStatus = "UNREACHABLE"
)

type PuzzleMetadata struct {
    Day  string `yaml:"day,omitempty"` 
    Title string `yaml:"title,omitempty"`
    Year string `yaml:"year,omitempty"`
}

type PuzzlePart struct {
    Answer string `yaml:"answer,omitempty"`
    Description string `yaml:"description,omitempty"`
    Status PuzzleStatus `yaml:"status,omitempty"`

    RawInput []byte
}

type Puzzle struct {
    Metadata PuzzleMetadata `yaml:"metadata"`
    Puzzles []PuzzlePart `yaml:"puzzles"`
}

func NewPuzzleFromHTML(day string, year string, htmlString string, input []byte) Puzzle {
    return Puzzle{
        Metadata: PuzzleMetadata{
            Day: day,
            Year: year,
            Title: getTitleFromBody(htmlString),
        },
        Puzzles: []PuzzlePart{},
    }
}

func NewPuzzleFromCache(filepath string, inputFilepath []string) Puzzle {
    var puzzle Puzzle
    yamlFile, err := os.ReadFile(filepath)
    if err != nil {
        log.Printf("Error trying to read the YAML file err =  #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &puzzle)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
   
    for i, inputFile := range inputFilepath {
        rawInput, err := os.ReadFile(inputFile)
        if err != nil {
            log.Printf("Error trying to read the input for Puzzle Part %d err   #%v ", i, err)
        }
        // we need to delete the last byte of the input because it is a newline or EOF
        rawInput = rawInput[:len(rawInput)-1]
        puzzle.Puzzles[i].RawInput = rawInput
    }
    return puzzle
}

func getTitleFromBody(body string) string {
    return "Title of the Puzzle Part"
}

type PuzzleSolver[T any] struct {
    Puzzle Puzzle
    NormalizeInput func(string) T
    Solve func() Response
}

type Response struct {
    Value string
    Error error
}
