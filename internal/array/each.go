// Package slices provides utility functions for working with slices.
package array

// Each iterates over a slice and calls a function (fn) for each element
func Each[T any](slice []T, fn func(idx int, element T)) {
	for i, item := range slice {
		fn(i, item)
	}
}
