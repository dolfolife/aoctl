package main

import (
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	beams := make(map[int]bool)
	splits := 0

	for x, char := range lines[0] {
		if char == 'S' {
			beams[x] = true
			break
		}
	}

	for y := 1; y < len(lines); y++ {
		nextBeams := make(map[int]bool)
		line := lines[y]
		for beamX := range beams {
			if beamX < 0 || beamX >= len(line) {
				continue
			}
			char := line[beamX]
			if char == '^' {
				splits++
				if beamX > 0 {
					nextBeams[beamX-1] = true
				}
				if beamX+1 < len(line) {
					nextBeams[beamX+1] = true
				}
			} else {
				nextBeams[beamX] = true
			}
		}
		beams = nextBeams
	}

	return strconv.Itoa(splits), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	beams := make(map[int]int)
	timelines := 0

	for x, char := range lines[0] {
		if char == 'S' {
			beams[x] = 1
			break
		}
	}

	for y := 1; y < len(lines); y++ {
		nextBeams := make(map[int]int)
		line := lines[y]
		for x, count := range beams {
			if x < 0 || x >= len(line) {
				timelines += count
				continue
			}
			char := line[x]
			if char == '^' {
				if x-1 >= 0 && x-1 < len(line) {
					nextBeams[x-1] += count
				} else {
					timelines += count
				}

				if x+1 >= 0 && x+1 < len(line) {
					nextBeams[x+1] += count
				} else {
					timelines += count
				}
			} else {
				nextBeams[x] += count
			}
		}
		beams = nextBeams
	}

	for _, count := range beams {
		timelines += count
	}

	return strconv.Itoa(timelines), nil
}
