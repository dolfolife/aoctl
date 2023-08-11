package main

import (
	"fmt"
	"io/ioutil"
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

func part1(bags []int) int {

	maxb := 0
	for i := 0; i < len(bags); i++ {
		if maxb < bags[i] {
			maxb = bags[i]
		}
	}

	return maxb
}

func part2(bags []int) int {

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

	return (maxb1 + maxb2 + maxb3)
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error processing the file")
	}

	e := strings.Split(string(b), "\n\n")
	bags, err := gmap(e[:len(e)-1], fromBagToCalories)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("part 1 => ", part1(bags))
	fmt.Println("part 2 => ", part2(bags))
}
