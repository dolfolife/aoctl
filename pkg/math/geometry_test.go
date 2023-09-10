package math

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCalDistancePoints(t *testing.T) {
    assert.Equal(t, "5", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: 5, Y: 0}))
    assert.Equal(t, "5", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: 0, Y: 5}))
    assert.Equal(t, "5", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: -5, Y: 0}))
    assert.Equal(t, "5", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: 0, Y: -5}))
    assert.Equal(t, "10", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: 5, Y: 5}))
    assert.Equal(t, "10", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: -5, Y: 5}))
    assert.Equal(t, "10", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: 5, Y: -5}))
    assert.Equal(t, "10", CalDistancePoints(Point{X: 0, Y: 0}, Point{X: -5, Y: -5}))
}

func TestCalIntersectionPoint(t *testing.T) {
    intersect, value := CalIntersectionPoint(Point{X: 0, Y: 0}, Point{X: 8, Y: 0}, Point{X: 4, Y: -4}, Point{X: 4, Y: 4})
    assert.Equal(t, true, intersect)
    assert.Equal(t, Point{4,0}, value)

    intersect, value = CalIntersectionPoint(Point{X: 4, Y: -4}, Point{X: 4, Y: 4}, Point{X: 0, Y: 0}, Point{X: 8, Y: 0})
    assert.Equal(t, true, intersect)
    assert.Equal(t, Point{4,0}, value)
    
    intersect, value = CalIntersectionPoint(Point{X: 0, Y: -8}, Point{X: 0, Y: -1}, Point{X: 0, Y: 0}, Point{X: 8, Y: 0})
    assert.Equal(t, false, intersect)
}
