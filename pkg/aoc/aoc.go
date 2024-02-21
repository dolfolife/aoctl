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
		log.Fatalf("Error creating directory at %s: %s", fullPath, err)
		return err
	}

	if err := createFileFromTemplates(filepath.Join(path, ".env"), GetRootFile("env")); err != nil {
		log.Fatalf("Error writing the env file: %s", err)
		return err
	}

	if err = createFileFromTemplates(filepath.Join(path, "README.md"), GetRootFile("README.md")); err != nil {
		log.Fatalf("Error creating README.md file: %s", err)
		return err
	}

	if err = createFileFromTemplates(filepath.Join(path, ".gitignore"), GetRootFile("gitignore")); err != nil {
		log.Fatalf("Error creating gitignore file: %s", err)
		return err
	}

	log.Println("Project initialized!")
	return nil
}

func createFileFromTemplates(file string, content string) error {
	roFile, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer roFile.Close()

	if _, err = roFile.WriteString(content); err != nil {
		return err
	}

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
