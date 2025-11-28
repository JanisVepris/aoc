package maths

// MinInt returns the smaller of two integers.
func MinInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
