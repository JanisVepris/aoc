package maths

import (
	"math"

	"janisvepris/aoc/internal/points"
)

// RectilinearDistance calculates the rectilinear distance between two 2D points.
// Also known as Manhattan distance or Taxicab distance
func RectilinearDistance(a, b points.Point2D) int {
	return AbsInt(a.X-b.X) + AbsInt(a.Y-b.Y)
}

// Distance2D calculates the Euclidean distance between two 2D points.
func Distance2D(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Hypot(dx, dy)
}

// Distance2DInt calculates the Euclidean distance between two 2D points and returns it as an integer.
func Distance2DInt(x1, y1, x2, y2 int) int {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)

	return int(math.Round(math.Hypot(dx, dy)))
}

// Distance3D calculates the Euclidean distance between two 3D points.
func Distance3D(x1, y1, z1, x2, y2, z2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	dz := z2 - z1

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Distance3DInt calculates the Euclidean distance between two 3D points and returns it as an integer.
func Distance3DInt(x1, y1, z1, x2, y2, z2 int) int {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	dz := float64(z2 - z1)

	return int(math.Round(math.Sqrt(dx*dx + dy*dy + dz*dz)))
}
