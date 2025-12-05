package array

// Push adds one or more elements to the end of the slice in place. The original slice is mutated.
func Push[T any](slice *[]T, elem ...T) {
	*slice = append(*slice, elem...)
}
