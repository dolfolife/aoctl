package main

import (
	"fmt"
	"io/ioutil"
)

func FindFirstMarker(distinct int, data []byte) int {
	// Keep track of the characters that have been seen
	var seen [26]bool

	// Iterate through the data, starting at the distinct-th character
	for i := distinct - 1; i < len(data); i++ {
		// Reset the seen array
		seen = [26]bool{}

		// Check if the distinct characters are all different
		different := true
		for j := i - distinct + 1; j <= i; j++ {
			// Convert the character to an index in the range 0-25
			index := data[j] - 'a'
			if seen[index] {
				different = false
				break
			}
			seen[index] = true
		}

		// If the last distinct characters are all different, return the index
		if different {
			return i + 1
		}
	}

	// If no marker was found, return -1
	return -1
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error processing the file")
	}
	fmt.Printf("Problem 1 %d, \n Problem 2 %d \n", FindFirstMarker(4, b), FindFirstMarker(14, b))
}
