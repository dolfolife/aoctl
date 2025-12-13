package main

import (
	"fmt"
	"strings"
)

func parseGraph(input string) map[string][]string {
	adj := make(map[string][]string)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {

		parts := strings.Split(line, ":")

		from := strings.TrimSpace(parts[0])
		dests := strings.Fields(parts[1])
		adj[from] = dests
	}
	return adj
}

func countPaths(adj map[string][]string, start, end string) int {
	memo := make(map[string]int)
	var dfs func(node string) int
	dfs = func(node string) int {
		if node == end {
			return 1
		}
		if val, ok := memo[node]; ok {
			return val
		}

		total := 0
		if neighbors, ok := adj[node]; ok {
			for _, neighbor := range neighbors {
				total += dfs(neighbor)
			}
		}

		memo[node] = total
		return total
	}
	return dfs(start)
}

func Part1(input string) (string, error) {
	adj := parseGraph(input)
	result := countPaths(adj, "you", "out")
	return fmt.Sprintf("%d", result), nil
}

func Part2(input string) (string, error) {
	adj := parseGraph(input)

	path1 := countPaths(adj, "svr", "dac") * countPaths(adj, "dac", "fft") * countPaths(adj, "fft", "out")

	path2 := countPaths(adj, "svr", "fft") * countPaths(adj, "fft", "dac") * countPaths(adj, "dac", "out")

	return fmt.Sprintf("%d", path1+path2), nil
}
