package aoc

import (
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

func InitializeProject(path string) error {
	fullPath := filepath.Join(os.Getenv("PWD"), path)

	log.Printf("....Initializing Advent of Code project in %s....", fullPath)

	err := os.Mkdir(fullPath, os.ModePerm)

	if err != nil {
		log.Fatalf("Error creating directory: %s", err)
	}

	cookieFile := filepath.Join(path, ".env")
	roFile, err := os.OpenFile(cookieFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error creating cookie file: %s", err)
	}

	defer roFile.Close()

	_, err = roFile.WriteString(`
# Use the .aoc file to set the environment variables for your project
AOC_SESSION=<insert-advent-of-code-session>
`)

	if err != nil {
		log.Fatalf("Error writing to the aoc file: %s", err)
	}

	readmeFile := filepath.Join(path, "README.md")
	roFile, err = os.OpenFile(readmeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error creating README.md file: %s", err)
	}

	defer roFile.Close()
	roFile.WriteString(`
    # Advent of Code
    `)

	gitIgnoreFile := filepath.Join(path, ".gitignore")
	roFile, err = os.OpenFile(gitIgnoreFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error creating .gitignore file: %s", err)
	}
	defer roFile.Close()

	if _, err = roFile.WriteString(`
# Advent of Code use the .aoc file for storing session information which you should
# not track in source control.
.aoc
`); err != nil {
		return err
	}

	log.Println("Project initialized")
	return nil
}

func GetPuzzles(day string, year string) []puzzle.Puzzle {

	aocConfig := GetAoCConfig()
	puzzleURL, err := url.JoinPath("https://adventofcode.com/", year, "/day/", day)

	if err != nil {
		log.Fatalf("Error creating url: %s", err)
	}

	body := getBodyFromUrl(puzzleURL, aocConfig.SessionId)

	inputURL, err := url.JoinPath(puzzleURL, "/input")

	if err != nil {
		log.Fatalf("Error creating url: %s", err)
	}

	rawInput := getBodyFromUrl(inputURL, aocConfig.SessionId)

	response, err := ParsePuzzles(day, year, body, rawInput)
	if err != nil {
		log.Fatalf("Error parsing puzzles: %s", err)
	}
	return response
}
