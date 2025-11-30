package main

import (
	"fmt"
	"strconv"
	"strings"
)

func gmap[T any, M any](data []T, f func(T) (M, error)) ([]M, error) {
	mapped := make([]M, len(data))
	for i, e := range data {
		v, err := f(e)
		if err != nil {
			return nil, err
		}
		mapped[i] = v
	}
	return mapped, nil
}

func fromBagToCalories(bag string) (int, error) {
	bc, err := gmap(strings.Split(bag, "\n"), strconv.Atoi)
	if err != nil {
		return 0, err
	}
	s := 0
	for _, a := range bc {
		s = s + a
	}
	return s, nil
}

func Part1(input string) (string, error) {
	e := strings.Split(input, "\n\n")
	// The original code did e[:len(e)-1] which suggests the last element might be empty due to trailing newline
	// We should check if the last element is empty
	if len(e) > 0 && strings.TrimSpace(e[len(e)-1]) == "" {
		e = e[:len(e)-1]
	}

	bags, err := gmap(e, fromBagToCalories)
	if err != nil {
		return "", err
	}

	maxb := 0
	for i := 0; i < len(bags); i++ {
		if maxb < bags[i] {
			maxb = bags[i]
		}
	}
	return fmt.Sprint(maxb), nil
}

func Part2(input string) (string, error) {
	e := strings.Split(input, "\n\n")
	if len(e) > 0 && strings.TrimSpace(e[len(e)-1]) == "" {
		e = e[:len(e)-1]
	}

	bags, err := gmap(e, fromBagToCalories)
	if err != nil {
		return "", err
	}

	maxb1, maxb2, maxb3 := 0, 0, 0
	for i := 0; i < len(bags); i++ {
		if bags[i] > maxb1 {
			maxb3 = maxb2
			maxb2 = maxb1
			maxb1 = bags[i]
		} else if bags[i] > maxb2 {
			maxb3 = maxb2
			maxb2 = bags[i]
		} else if bags[i] > maxb3 {
			maxb3 = bags[i]
		}
	}
	return fmt.Sprint(maxb1 + maxb2 + maxb3), nil
}
