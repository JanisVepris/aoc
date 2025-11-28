// Package maths provides basic mathematical functions.
package maths

// AbsInt returns the absolute value of an integer.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
