package aoc

import (
	"log"
	"math"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

	rootFiles := ProjectTemplates()

	for filename, content := range rootFiles {
		if err := createFileFromTemplates(filepath.Join(path, filename), content); err != nil {
			log.Fatalf("Error writing the %s file: %s\n", filename, err)
			return err
		}
	}

	log.Println("Project initialized!")
	return nil
}

func SyncAoC(force bool) {
	firstYearOfAoC := 2015
	currentYear, CurrentMonth, currentDay := time.Now().Date()

	// if we are not in december it means that we do not have AoC for the current year
	if CurrentMonth < 12 {
		currentYear--
	}
	config := GetAoCConfig()
	eventsPath := filepath.Join(config.ProjectPath, "events")

	err := os.Mkdir(eventsPath, os.ModePerm)

	if err != nil && !os.IsExist(err) {
		log.Fatal("events directory cannot be created")
	}

	for i := firstYearOfAoC; i <= currentYear; i++ {
		yearDirPath := filepath.Join(eventsPath, strconv.Itoa(i))

		err := os.Mkdir(yearDirPath, os.ModePerm)

		if err != nil && !os.IsExist(err) {
			log.Fatalf("events directory for year %d cannot be created\n", i)
		}

		// if we are in December we take the min day between Today and 25
		lastDay := 25
		if CurrentMonth == 12 {
			lastDay = int(math.Min(float64(lastDay), float64(currentDay)))
		}

		for j := 1; j <= lastDay; j++ {
			dayDirPath := filepath.Join(yearDirPath, strconv.Itoa(j))
			err := os.Mkdir(dayDirPath, os.ModePerm)

			if err != nil && !os.IsExist(err) {
				log.Fatalf("events directory for year %d cannot be created\n", i)
			}

			if os.IsExist(err) && !force {
				continue
			}
			// TODO:
			// create all files
			// - Create puzzleDay.yaml
			// - Create main.go
			// - Create solution.go (with two parts)
			// - create input directory
			// - create solution_test.go
			// - Create README.md (use for cache the text in the adventofcode.com
		}
	}
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
