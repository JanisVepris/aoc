package maths

import "janisvepris/aoc25/internal/points"

// RectilinearDistance calculates the rectilinear distance between two 2D points.
// Also known as Manhattan distance or Taxicab distance
func RectilinearDistance(a, b points.Point2D) int {
	return AbsInt(a.X-b.X) + AbsInt(a.Y-b.Y)
}
