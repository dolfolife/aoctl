package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	totalPresses := 0

	for _, line := range lines {
		target, buttons := parseLine(line)
		minPresses := solveMachine(target, buttons)
		totalPresses += minPresses
	}

	return fmt.Sprintf("%d", totalPresses), nil
}

func parseLine(line string) ([]int, [][]int) {

	line = strings.TrimSpace(line)

	targetStr := line[strings.Index(line, "[")+1 : strings.Index(line, "]")]
	var target []int
	for _, c := range targetStr {
		if c == '#' {
			target = append(target, 1)
		} else {
			target = append(target, 0)
		}
	}

	re := regexp.MustCompile(`\(([\d,]+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	var buttons [][]int
	numLights := len(target)

	for _, match := range matches {
		content := match[1]
		parts := strings.Split(content, ",")
		btn := make([]int, numLights)
		for _, part := range parts {
			idx, _ := strconv.Atoi(part)
			if idx < numLights {
				btn[idx] = 1
			}
		}
		buttons = append(buttons, btn)
	}

	return target, buttons
}

func solveMachine(target []int, buttons [][]int) int {
	numRows := len(target)
	numCols := len(buttons)

	matrix := make([][]int, numRows)
	for r := 0; r < numRows; r++ {
		matrix[r] = make([]int, numCols+1)
		for c := 0; c < numCols; c++ {
			matrix[r][c] = buttons[c][r]
		}
		matrix[r][numCols] = target[r]
	}

	pivotRow := 0
	pivotCols := make([]int, 0)
	freeCols := make([]int, 0)

	isPivot := make(map[int]bool)

	for c := 0; c < numCols && pivotRow < numRows; c++ {
		sel := -1
		for r := pivotRow; r < numRows; r++ {
			if matrix[r][c] == 1 {
				sel = r
				break
			}
		}

		if sel == -1 {
			continue
		}

		matrix[pivotRow], matrix[sel] = matrix[sel], matrix[pivotRow]

		for r := 0; r < numRows; r++ {
			if r != pivotRow && matrix[r][c] == 1 {
				for k := c; k <= numCols; k++ {
					matrix[r][k] ^= matrix[pivotRow][k]
				}
			}
		}

		pivotCols = append(pivotCols, c)
		isPivot[c] = true
		pivotRow++
	}

	for r := pivotRow; r < numRows; r++ {
		if matrix[r][numCols] == 1 {
			return 0
		}
	}

	for c := 0; c < numCols; c++ {
		if !isPivot[c] {
			freeCols = append(freeCols, c)
		}
	}

	xp := make([]int, numCols)
	for i := 0; i < len(pivotCols); i++ {
		xp[pivotCols[i]] = matrix[i][numCols]
	}

	var nullBasis [][]int
	for _, f := range freeCols {
		basis := make([]int, numCols)
		basis[f] = 1
		for i := 0; i < len(pivotCols); i++ {
			p := pivotCols[i]
			if matrix[i][f] == 1 {
				basis[p] = 1
			}
		}
		nullBasis = append(nullBasis, basis)
	}

	minWeight := math.MaxInt32

	numFree := len(freeCols)
	limit := 1 << numFree

	for i := 0; i < limit; i++ {
		sol := make([]int, numCols)
		copy(sol, xp)

		for b := 0; b < numFree; b++ {
			if (i>>b)&1 == 1 {
				for k := 0; k < numCols; k++ {
					sol[k] ^= nullBasis[b][k]
				}
			}
		}

		weight := 0
		for _, v := range sol {
			weight += v
		}
		if weight < minWeight {
			minWeight = weight
		}
	}

	return minWeight
}

func Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	totalPresses := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		target, buttons := parseLinePart2(line)
		minPresses := solveMachineInt(target, buttons)
		if minPresses == -1 {
			return "", fmt.Errorf("no solution found for machine: %s", line)
		}
		totalPresses += minPresses
	}

	return fmt.Sprintf("%d", totalPresses), nil
}

func parseLinePart2(line string) ([]int, [][]int) {
	line = strings.TrimSpace(line)
	targetStr := line[strings.Index(line, "{")+1 : strings.Index(line, "}")]
	targetParts := strings.Split(targetStr, ",")
	var target []int
	for _, p := range targetParts {
		val, _ := strconv.Atoi(p)
		target = append(target, val)
	}

	re := regexp.MustCompile(`\(([\d,]+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	var buttons [][]int
	numCounters := len(target)

	for _, match := range matches {
		content := match[1]
		parts := strings.Split(content, ",")
		btn := make([]int, numCounters)
		for _, part := range parts {
			idx, _ := strconv.Atoi(part)
			if idx < numCounters {
				btn[idx] = 1
			}
		}
		buttons = append(buttons, btn)
	}

	return target, buttons
}

func solveMachineInt(target []int, buttons [][]int) int {
	numRows := len(target)
	numCols := len(buttons)

	matrix := make([][]float64, numRows)
	for r := 0; r < numRows; r++ {
		matrix[r] = make([]float64, numCols+1)
		for c := 0; c < numCols; c++ {
			matrix[r][c] = float64(buttons[c][r])
		}
		matrix[r][numCols] = float64(target[r])
	}

	pivotRow := 0
	pivotCols := make([]int, 0)
	freeCols := make([]int, 0)
	isPivot := make(map[int]bool)

	const epsilon = 1e-9

	for c := 0; c < numCols && pivotRow < numRows; c++ {
		sel := -1
		for r := pivotRow; r < numRows; r++ {
			if math.Abs(matrix[r][c]) > epsilon {
				sel = r
				break
			}
		}
		if sel == -1 {
			continue
		}

		matrix[pivotRow], matrix[sel] = matrix[sel], matrix[pivotRow]

		pivotVal := matrix[pivotRow][c]
		for k := c; k <= numCols; k++ {
			matrix[pivotRow][k] /= pivotVal
		}

		for r := 0; r < numRows; r++ {
			if r != pivotRow {
				factor := matrix[r][c]
				if math.Abs(factor) > epsilon {
					for k := c; k <= numCols; k++ {
						matrix[r][k] -= factor * matrix[pivotRow][k]
					}
				}
			}
		}

		pivotCols = append(pivotCols, c)
		isPivot[c] = true
		pivotRow++
	}

	for r := pivotRow; r < numRows; r++ {
		if math.Abs(matrix[r][numCols]) > epsilon {
			return 0
		}
	}

	for c := 0; c < numCols; c++ {
		if !isPivot[c] {
			freeCols = append(freeCols, c)
		}
	}

	xp := make([]float64, numCols)
	for i := 0; i < len(pivotCols); i++ {
		xp[pivotCols[i]] = matrix[i][numCols]
	}
	var nullBasis [][]float64
	for _, f := range freeCols {
		basis := make([]float64, numCols)
		basis[f] = 1.0
		for i := 0; i < len(pivotCols); i++ {
			p := pivotCols[i]
			basis[p] = -matrix[i][f]
		}
		nullBasis = append(nullBasis, basis)
	}

	result := math.MaxInt32
	found := false

	var search func(idx int, currentSol []float64)
	search = func(idx int, currentSol []float64) {
		allNonNeg := true
		allInt := true

		weight := 0
		for _, v := range currentSol {
			if v < -epsilon {
				allNonNeg = false
				break
			}
			if math.Abs(v-math.Round(v)) > epsilon {
				allInt = false
			}
			weight += int(math.Round(v))
		}

		if allNonNeg && allInt {
			if weight < result {
				result = weight
				found = true
			}
		}

		if idx == len(freeCols) {
			return
		}

		basis := nullBasis[idx]
		limit := 100

		for k := 0; k <= limit; k++ {
			nextSol := make([]float64, len(currentSol))
			copy(nextSol, currentSol)

			search(idx+1, nextSol)

			for i := 0; i < len(currentSol); i++ {
				currentSol[i] += basis[i]
			}

		}
		for i := 0; i < len(currentSol); i++ {
			currentSol[i] -= float64(limit+1) * basis[i]
		}
	}

	var recursiveSearch func(idx int, c []int)
	recursiveSearch = func(idx int, c []int) {
		if idx == len(freeCols) {
			currentSol := make([]float64, numCols)
			copy(currentSol, xp)
			valid := true
			weight := 0

			for i, val := range c {
				for k := 0; k < numCols; k++ {
					currentSol[k] += float64(val) * nullBasis[i][k]
				}
			}

			for _, v := range currentSol {
				if v < -epsilon {
					valid = false
					break
				}
				if math.Abs(v-math.Round(v)) > epsilon {
					valid = false
					break
				}
				weight += int(math.Round(v))
			}

			if valid {
				if weight < result {
					result = weight
					found = true
				}
			}
			return
		}

		for k := 0; k <= 200; k++ {
			c[idx] = k
			recursiveSearch(idx+1, c)
		}
	}

	coeffs := make([]int, len(freeCols))
	recursiveSearch(0, coeffs)

	if !found {
		return -1
	}
	return result
}
