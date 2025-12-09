package main

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var Part1Limit = 1000

type Point struct {
	X int
	Y int
	Z int
}

type Edge struct {
	P1       int
	P2       int
	Distance float64
}

type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type UnionFind struct {
	parent []int
	size   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size, count: n}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.Find(uf.parent[i])
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) bool {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI != rootJ {
		if uf.size[rootI] < uf.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		uf.parent[rootJ] = rootI
		uf.size[rootI] += uf.size[rootJ]
		uf.count--
		return true
	}
	return false
}

func pythagoreanDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2) + math.Pow(float64(p1.Z-p2.Z), 2))
}

func Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var points []Point
	for _, line := range lines {
		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		points = append(points, Point{X: x, Y: y, Z: z})
	}

	if len(points) == 0 {
		return "", errors.New("no points found")
	}

	h := &EdgeHeap{}
	heap.Init(h)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := pythagoreanDistance(points[i], points[j])
			heap.Push(h, Edge{P1: i, P2: j, Distance: dist})
		}
	}

	uf := NewUnionFind(len(points))
	connectionsMade := 0

	for h.Len() > 0 && connectionsMade < Part1Limit {
		edge := heap.Pop(h).(Edge)

		uf.Union(edge.P1, edge.P2)
		connectionsMade++
	}

	circuitSizes := make(map[int]int)
	for i := 0; i < len(points); i++ {
		root := uf.Find(i)
		circuitSizes[root] = uf.size[root]
	}

	var sizes []int
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= sizes[i]
	}

	return fmt.Sprintf("%d", result), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var points []Point
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		points = append(points, Point{X: x, Y: y, Z: z})
	}

	if len(points) == 0 {
		return "", errors.New("no points found")
	}

	h := &EdgeHeap{}
	heap.Init(h)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := pythagoreanDistance(points[i], points[j])
			heap.Push(h, Edge{P1: i, P2: j, Distance: dist})
		}
	}

	uf := NewUnionFind(len(points))

	for h.Len() > 0 {
		edge := heap.Pop(h).(Edge)
		if uf.Union(edge.P1, edge.P2) {
			if uf.count == 1 {
				return fmt.Sprintf("%d", points[edge.P1].X*points[edge.P2].X), nil
			}
		}
	}

	return "", errors.New("could not connect all points")
}
