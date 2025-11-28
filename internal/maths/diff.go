package maths

import "math"

// DiffInt returns the absolute difference between two integers
func DiffInt(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
