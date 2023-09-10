package math

import (
    "math"
    "strconv"
)

type Point struct {
    X, Y float64
}

type Segment struct {
    A, B Point
}

func CalDistancePoints(pos1 Point, pos2 Point) string {
    partX := math.Abs(float64(pos1.X) - float64(pos2.X))
    partY := math.Abs(float64(pos1.Y) - float64(pos2.Y))
    return strconv.Itoa(int(partX + partY))
}


func CalIntersectionPoint(a, b, c, d Point) (bool, Point) {
    cross := (c.Y-d.Y)*(b.X-a.X) - (c.X-d.X)*(b.Y-a.Y)
    if cross == 0 {
        return false, Point{} // Parallel lines, no intersection
    }

    t1 := float64((c.Y-d.Y)*(c.X-a.X) + (d.X-c.X)*(c.Y-a.Y)) / float64(cross)
    t2 := float64((a.Y-b.Y)*(c.X-a.X) + (b.X-a.X)*(c.Y-a.Y)) / float64(cross)

    p1 := Point{b.X - a.X, b.Y - a.Y}
    // Check if the intersection point is within the line segments
    if t1 >= 0 && t1 <= 1 && t2 >= 0 && t2 <= 1 {
        intersectionX := a.X + t1*p1.X
        intersectionY := a.Y + t1*p1.Y
        return true, Point{intersectionX, intersectionY}
    }

    return false, Point{} // The intersection point is outside the line segments
}

func IsValidTriangle(a, b, c int) bool {
    return a+b > c && a+c > b && b+c > a
}
