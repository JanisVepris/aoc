package array

// Reduce applies a function of two arguments cumulatively to the elements
//
// Reduce iterates over a slice and calls a function (fn) for each element
// The function takes a carry and an element and returns a new carry
// The initial carry is passed as the third argument
// The final carry is returned
func Reduce[T any, U any](slice []T, fn func(carry U, element T) U, initial U) U {
	result := initial

	for _, item := range slice {
		result = fn(result, item)
	}

	return result
}
