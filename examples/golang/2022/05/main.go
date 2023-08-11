package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Crate []string

func (c *Crate) IsEmpty() bool {
	return len(*c) == 0
}

func (c *Crate) Push(e []string) {
	*c = append(*c, e...)
}

func (c *Crate) Pop(i int) []string {
	if c.IsEmpty() {
		return nil
	}

	index := len(*c) - 1
	e := (*c)[index-i+1:]
	*c = (*c)[:index-i+1]
	return e
}

func findCratesForIndex(m []string, pivot int) Crate {
	c := Crate{}
	for i := len(m) - 1; i >= 0; i-- {
		row := strings.Split(m[i], "")
		if row[pivot] != " " {
			c.Push([]string{row[pivot]})
		}
	}
	return c
}

func buildCrates(b string) []Crate {

	s := strings.Split(b, "\n")
	lastRow := strings.Split(s[len(s)-1], "")
	crates := make([]Crate, 0)

	for i := 0; i < len(lastRow); i++ {

		if lastRow[i] != " " {
			crates = append(crates, findCratesForIndex(s[:len(s)-1], i))
		}
	}

	return crates
}

func playMoves(c []Crate, m []string, byOne bool) []Crate {

	re := regexp.MustCompile("[0-9]+")
	for i := 0; i < len(m)-1; i++ {
		commandLine := re.FindAllString(m[i], -1)
		numOfMoves, _ := strconv.Atoi(commandLine[0])
		from, _ := strconv.Atoi(commandLine[1])
		to, _ := strconv.Atoi(commandLine[2])
		if byOne {
			for j := 0; j < numOfMoves; j++ {
				e := c[from-1].Pop(1)
				c[to-1].Push(e)
			}
		} else {
			e := c[from-1].Pop(numOfMoves)
			c[to-1].Push(e)
		}
	}
	return c
}

func main() {

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error processing the file")
	}

	input := strings.Split(string(b), "\n\n")

	finalCratesPart1 := playMoves(buildCrates(input[0]), strings.Split(input[1], "\n"), true)
	finalCratesPart2 := playMoves(buildCrates(input[0]), strings.Split(input[1], "\n"), false)

	topBlock1 := ""
	topBlock2 := ""
	for i := 0; i < len(finalCratesPart1); i++ {
		topBlock1 = topBlock1 + finalCratesPart1[i].Pop(1)[0]
		topBlock2 = topBlock2 + finalCratesPart2[i].Pop(1)[0]
	}

	fmt.Printf("Solutions: \n Part1 = %s \n Part2 = %s \n", topBlock1, topBlock2)
}
