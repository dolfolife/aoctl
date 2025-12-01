package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Move struct {
	Direction byte
	Steps     int
}

func parseLine(line string) (Move, error) {

	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return Move{}, nil
	}

	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		return Move{}, err
	}

	return Move{
		Direction: line[0],
		Steps:     steps,
	}, nil
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	timesZero := 0
	arrow := 50
	for _, line := range lines {
		move, err := parseLine(line)
		if err != nil {
			return "", err
		}

		if move.Direction == 'L' {
			arrow -= move.Steps % 100
		} else {
			arrow += move.Steps % 100
		}

		if arrow < 0 {
			arrow = 100 + arrow
		}
		if arrow > 100 {
			arrow = arrow - 100
		}
		if arrow == 0 || arrow == 100 {
			timesZero++
			arrow = 0
		}
	}
	return fmt.Sprint(timesZero), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	timesZero := 0
	arrow := 50
	for _, line := range lines {
		move, err := parseLine(line)
		if err != nil {
			return "", err
		}

		if move.Direction == 'R' {
			timesZero += int(math.Floor(float64(arrow+move.Steps)/100.0) - math.Floor(float64(arrow)/100.0))
			arrow = (arrow + move.Steps) % 100
		} else {
			timesZero += int(math.Floor(float64(arrow-1)/100.0) - math.Floor(float64(arrow-move.Steps-1)/100.0))
			arrow = (arrow - move.Steps) % 100
			if arrow < 0 {
				arrow += 100
			}
		}
	}
	return fmt.Sprint(timesZero), nil
}
