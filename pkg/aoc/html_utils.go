package aoc

import (
	"bytes"
	"log"
	"strings"

	"github.com/dolfolife/aoctl/pkg/puzzle"
	"golang.org/x/net/html"
)

func ParsePuzzles(day string, year string, responseBody []byte, input []byte) ([]puzzle.Puzzle, error) {
	node, err := html.Parse(bytes.NewReader(responseBody))

	if err != nil {
		log.Fatalf("Error parsing the body: %s", err)
	}

	puzzlePartsHTMLNodes := findRootNodesPuzzle(node)

	var parts []puzzle.Puzzle
	for _, node := range puzzlePartsHTMLNodes {
		parts = append(parts, puzzle.NewPuzzleFromHTML(day, year, parsePuzzleHTML(node), input))
	}
	return parts, nil
}

func findRootNodesPuzzle(node *html.Node) []*html.Node {
	var nodes []*html.Node

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode && child.Data == "article" && hasAttr(child.Attr, "day-desc") {
			nodes = append(nodes, child)
		}
		nodes = append(nodes, findRootNodesPuzzle(child)...)
	}

	return nodes
}

func hasAttr(attrs []html.Attribute, attr string) bool {
	for _, a := range attrs {
		if a.Key == "class" && a.Val == attr {
			return true
		}
	}
	return false
}

func parsePuzzleHTML(node *html.Node) string {
	buffer := strings.Builder{}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			buffer.WriteString(child.Data)
		}
		buffer.WriteString(parsePuzzleHTML(child))
	}

	return buffer.String()
}
