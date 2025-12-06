package main

import (
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	From int
	To   int
}

func Part1(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	ranges, ids := parts[0], parts[1]

	rangesMap := make(map[Range]bool)
	for _, r := range strings.Split(ranges, "\n") {
		freshIds := strings.Split(r, "-")
		first, _ := strconv.Atoi(freshIds[0])
		last, _ := strconv.Atoi(freshIds[1])
		rangesMap[Range{From: first, To: last}] = true
	}

	result := 0
	for _, id := range strings.Split(ids, "\n") {
		v, _ := strconv.Atoi(id)
		for r := range rangesMap {
			if v >= r.From && v <= r.To {
				result++
				break
			}
		}
	}
	return strconv.Itoa(result), nil
}

func Part2(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	rangesStr := parts[0]

	var ranges []Range
	for _, r := range strings.Split(rangesStr, "\n") {
		freshIds := strings.Split(r, "-")
		first, _ := strconv.Atoi(freshIds[0])
		last, _ := strconv.Atoi(freshIds[1])
		ranges = append(ranges, Range{From: first, To: last})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})

	var merged []Range
	if len(ranges) > 0 {
		current := ranges[0]
		for i := 1; i < len(ranges); i++ {
			next := ranges[i]
			if current.To >= next.From {
				if next.To > current.To {
					current.To = next.To
				}
			} else {
				merged = append(merged, current)
				current = next
			}
		}
		merged = append(merged, current)
	}

	result := 0
	for _, r := range merged {
		result += r.To - r.From + 1
	}

	return strconv.Itoa(result), nil
}
