package aoc

import (
    "bytes"
    "log"
    "strings"

    "golang.org/x/net/html"
)

func getPuzzleParts(body []byte) []string {
    node, err := html.Parse(bytes.NewReader(body))
   
    if err != nil {
        log.Fatalf("Error parsing the body: %s", err)
    }

    puzzlePartsHTMLNodes := findRootNodesPuzzlePart(node)
    
    var parts []string
    for _, node := range puzzlePartsHTMLNodes {
       parts = append(parts, parsePuzzleHTML(node))
    }
    return parts
}

func findRootNodesPuzzlePart(node *html.Node) []*html.Node {
    var nodes []*html.Node
    
    for child := node.FirstChild; child != nil; child = child.NextSibling {
        if child.Type == html.ElementNode && child.Data == "article" && hasAttr(child.Attr, "day-desc") {
            nodes = append(nodes, child)
        }
        nodes = append(nodes, findRootNodesPuzzlePart(child)...)
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
