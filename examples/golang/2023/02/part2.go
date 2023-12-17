package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dolfolife/aoctl/pkg/puzzle"
)

type Day2Part2Solver struct {
	Puzzle puzzle.PuzzlePart
}

func setMaxVal(color string, maxVal *map[string]int) {

}
func (s *Day2Part2Solver) NormalizeInput(input string) []int {
	result := []int{}
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`Game (?P<gameid>[0-9]+)\: (?P<colors>.*)`)

	for _, line := range lines {
		val := r.FindAllStringSubmatch(line, -1)
		colors := val[0][2]

		maxVals := map[string]int{"blue": 0, "red": 0, "green": 0}
		newC := strings.Split(colors, ";")
		for _, c := range newC {
			colorStringList := strings.Split(c, ",")
			for _, colorString := range colorStringList {
				colorString = strings.Trim(colorString, " ")
				colorString = strings.Trim(colorString, "\n")
				colorString = strings.Trim(colorString, "\t")
				colorDic := strings.Split(colorString, " ")
				bagColorNumber, _ := strconv.Atoi(colorDic[0])
				if maxVals[colorDic[1]] < bagColorNumber {
					maxVals[colorDic[1]] = bagColorNumber
				}
			}
		}
		result = append(result, maxVals["blue"]*maxVals["red"]*maxVals["green"])
	}
	return result
}
func (s *Day2Part2Solver) Solve() (string, error) {
	return part2(s.NormalizeInput(string(s.Puzzle.RawInput)))
}

func part2(input []int) (string, error) {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return fmt.Sprint(sum), nil
}
