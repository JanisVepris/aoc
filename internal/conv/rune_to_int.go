// Package conv provides functions for converting between types.
package conv

import "strconv"

// RuneToInt converts a rune representing a digit to its integer value.
func RuneToInt(value rune) int {
	result, err := strconv.Atoi(string(value))
	if err != nil {
		panic(err)
	}

	return result
}
