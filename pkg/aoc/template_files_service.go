package aoc

import (
	"embed"
	"fmt"
	"log"
)

//go:embed templates
var templates embed.FS

func GetRootFile(filename string) string {
	data, err := templates.ReadFile(fmt.Sprintf("templates/root/%s", filename))

	if err != nil {
		log.Fatalf("error reading %s template with: \n%v\n", filename, err)
	}

	return string(data)
}

func GetSolutionFile(filename string) string {
	data, err := templates.ReadFile(fmt.Sprintf("templates/solution/%s", filename))

	if err != nil {
		log.Fatalf("error reading %s template with: \n%v\n", filename, err)
	}

	return string(data)
}
