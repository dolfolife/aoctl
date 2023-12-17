package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

type Day2Part1Solver struct {
	Puzzle puzzle.PuzzlePart
}

func (s *Day2Part1Solver) NormalizeInput(input string) []int {
	result := []int{}
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`Game (?P<gameid>[0-9]+)\: (?P<colors>.*)`)

	for _, line := range lines {
		val := r.FindAllStringSubmatch(line, -1)
		gameId := val[0][1]
		colors := val[0][2]

		newC := strings.Split(colors, ";")
		isLineValid := true
		for _, c := range newC {
			colorStringList := strings.Split(c, ",")
			for _, colorString := range colorStringList {
				colorString = strings.Trim(colorString, " ")
				colorString = strings.Trim(colorString, "\n")
				colorString = strings.Trim(colorString, "\t")
				if !canExistColor(colorString) {
					isLineValid = false
					break
				}
			}
		}
		if isLineValid {
			i, _ := strconv.Atoi(gameId)
			result = append(result, i)
		}
	}
	return result
}

func canExistColor(color string) bool {
	possibleColorSize := map[string]int{"blue": 14, "red": 12, "green": 13}
	colorDic := strings.Split(color, " ")
	bagColorNumber, _ := strconv.Atoi(colorDic[0])
	return possibleColorSize[colorDic[1]] >= bagColorNumber
}

func (s *Day2Part1Solver) Solve() (string, error) {
	return part1(s.NormalizeInput(string(s.Puzzle.RawInput)))
}

func part1(input []int) (string, error) {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return fmt.Sprint(sum), nil
}
