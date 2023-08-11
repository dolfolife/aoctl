package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func calPos(r rune) int {
	pos := 0
	if unicode.IsUpper(r) {
		pos = pos + 26
	}

	return int(rune(r)&31) + pos
}

func findRepeatedChar(items []rune) int {
	repeatedItem := 0

	firstCompartment := items[:len(items)/2]
	secondCompartment := items[len(items)/2:]
	abc := make([]bool, 52)
	for i := 0; i < len(firstCompartment); i++ {
		abc[calPos(firstCompartment[i])-1] = true
	}
	for j := 0; j < len(secondCompartment); j++ {

		if abc[calPos(secondCompartment[j])-1] {
			repeatedItem = calPos(secondCompartment[j])
			break
		}
	}

	return repeatedItem
}

func solvePart1(rucksacks []string) int {

	total := 0
	for i := 0; i < len(rucksacks); i++ {

		total = total + findRepeatedChar([]rune(rucksacks[i]))
	}
	return total
}

func findBadge(bag1 []rune, bag2 []rune, bag3 []rune) int {
	repeatedItem := 0

	abc := make([]bool, 52)
	abc2 := make([]bool, 52)
	for i := 0; i < len(bag1); i++ {
		abc[calPos(bag1[i])-1] = true
	}
	for i := 0; i < len(bag2); i++ {
		if abc[calPos(bag2[i])-1] {
			abc2[calPos(bag2[i])-1] = true
		}
	}

	for j := 0; j < len(bag3); j++ {

		if abc2[calPos(bag3[j])-1] {
			repeatedItem = calPos(bag3[j])
		}
	}
	return repeatedItem
}

func solvePart2(rucksacks []string) int {

	total := 0
	for i := 0; i+3 <= len(rucksacks); i = i + 3 {
		total = total + findBadge([]rune(rucksacks[i]), []rune(rucksacks[i+1]), []rune(rucksacks[i+2]))
	}

	return total
}

func main() {

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error processing the file")
	}

	rucksacks := strings.Split(string(b), "\n")

	part1Solution := solvePart1(rucksacks)
	part2Solution := solvePart2(rucksacks)

	fmt.Printf("solutions: \n Part1 = %d \n Part2 = %d \n", part1Solution, part2Solution)
}
