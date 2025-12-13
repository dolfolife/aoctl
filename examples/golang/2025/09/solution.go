package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var points []Point

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		points = append(points, Point{x, y})
	}

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			width := int(math.Abs(float64(p1.x-p2.x))) + 1
			height := int(math.Abs(float64(p1.y-p2.y))) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return fmt.Sprintf("%d", maxArea), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var points []Point

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		points = append(points, Point{x, y})
	}

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			width := int(math.Abs(float64(p1.x-p2.x))) + 1
			height := int(math.Abs(float64(p1.y-p2.y))) + 1
			area := width * height

			if area <= maxArea {
				continue
			}

			if isValid(p1, p2, points) {
				maxArea = area
			}
		}
	}

	return fmt.Sprintf("%d", maxArea), nil
}

func isValid(p1, p2 Point, poly []Point) bool {
	minX := math.Min(float64(p1.x), float64(p2.x))
	maxX := math.Max(float64(p1.x), float64(p2.x))
	minY := math.Min(float64(p1.y), float64(p2.y))
	maxY := math.Max(float64(p1.y), float64(p2.y))

	cx := (minX + maxX) / 2.0
	cy := (minY + maxY) / 2.0
	if !isPointInPoly(cx, cy, poly) {
		return false
	}
	for k := 0; k < len(poly); k++ {
		v1 := poly[k]
		v2 := poly[(k+1)%len(poly)]

		if v1.x == v2.x {
			vx := float64(v1.x)
			if vx > minX && vx < maxX {
				vy1 := math.Min(float64(v1.y), float64(v2.y))
				vy2 := math.Max(float64(v1.y), float64(v2.y))
				overlapStart := math.Max(vy1, minY)
				overlapEnd := math.Min(vy2, maxY)
				if overlapStart < overlapEnd {
					return false
				}
			}
		} else {
			vy := float64(v1.y)
			if vy > minY && vy < maxY {
				vx1 := math.Min(float64(v1.x), float64(v2.x))
				vx2 := math.Max(float64(v1.x), float64(v2.x))
				overlapStart := math.Max(vx1, minX)
				overlapEnd := math.Min(vx2, maxX)
				if overlapStart < overlapEnd {
					return false
				}
			}
		}
	}

	return true
}

func isPointInPoly(x, y float64, poly []Point) bool {

	n := len(poly)
	inside := false

	for k := range n {
		v1 := poly[k]
		v2 := poly[(k+1)%n]
		if onSegment(x, y, v1, v2) {
			return true
		}
	}

	for k := range n {
		v1 := poly[k]
		v2 := poly[(k+1)%n]

		p1x, p1y := float64(v1.x), float64(v1.y)
		p2x, p2y := float64(v2.x), float64(v2.y)

		if (p1y > y) != (p2y > y) {
			xInt := p1x + (y-p1y)*(p2x-p1x)/(p2y-p1y)
			if x < xInt {
				inside = !inside
			}
		}
	}
	return inside
}

func onSegment(x, y float64, v1, v2 Point) bool {
	p1x, p1y := float64(v1.x), float64(v1.y)
	p2x, p2y := float64(v2.x), float64(v2.y)

	if p1x == p2x {
		return x == p1x && y >= math.Min(p1y, p2y) && y <= math.Max(p1y, p2y)
	}
	if p1y == p2y {
		return y == p1y && x >= math.Min(p1x, p2x) && x <= math.Max(p1x, p2x)
	}
	return false
}
