package array

// PopInPlace removes the last element from the slice in place and returns it. The original slice is mutated.
func PopInPlace[T any](slice *[]T) T {
	s := *slice
	if len(s) == 0 {
		var zero T
		return zero
	}
	elem := s[len(s)-1]
	*slice = s[:len(s)-1]
	return elem
}
