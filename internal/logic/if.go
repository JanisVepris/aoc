// Package logic provides utilities for logical operations.
package logic

// If returns trueVal if condition is true, otherwise it returns falseVal
func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}
