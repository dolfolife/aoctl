package aoc

import (
    "log"
    "os"
    "path/filepath"
    "net/url"

)

func InitializeProject(path string) error {
    fullPath := filepath.Join(os.Getenv("PWD"), path)
    
    log.Printf("....Initializing Advent of Code project in %s....", fullPath)

    err := os.Mkdir(fullPath, os.ModePerm)

    if err != nil {
        log.Fatalf("Error creating directory: %s", err)
    }
    
    cookieFile := filepath.Join(path, ".aoc")
    roFile, err := os.OpenFile(cookieFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        log.Fatalf("Error creating cookie file: %s", err)
    }
    
    defer roFile.Close()

    _, err = roFile.WriteString(`
# Use the .aoc file to set the environment variables for your project
session=<insert-advent-of-code-session>
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


func GetPuzzles(day string, year string, cookie string) []string {
    url, err := url.JoinPath("https://adventofcode.com/", year, "/day/", day)
    
    if err != nil {
        log.Fatalf("Error creating url: %s", err)
    }
    
    body := getBodyFromUrl(url, cookie)

    return getPuzzleParts(body)
}
