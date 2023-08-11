package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// a1-a2, b1-b2
func solvePart1(a1 int, a2 int, b1 int, b2 int) bool {
	return (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2)
}

func solvePart2(a1 int, a2 int, b1 int, b2 int) bool {
	return a2 >= b1 && b2 >= a1
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func solveSections(assigments []string) (int, int) {

	total1 := 0
	total2 := 0
	for i := 0; i < len(assigments)-1; i++ {
		sections := strings.Split(assigments[i], ",")
		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")

		if solvePart1(atoi(section1[0]), atoi(section1[1]), atoi(section2[0]), atoi(section2[1])) {
			total1++
		}
		if solvePart2(atoi(section1[0]), atoi(section1[1]), atoi(section2[0]), atoi(section2[1])) {
			total2++
		}
	}

	return total1, total2
}

func main() {

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error processing the file")
	}

	assignments := strings.Split(string(b), "\n")

	part1Solution, part2Solution := solveSections(assignments)

	fmt.Printf("solutions: \n Part1 = %d \n Part2 = %d \n", part1Solution, part2Solution)
}
