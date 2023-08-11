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
    // Create .coe cookie file
    // Create README.md from a template
    // Create .gitignore

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
