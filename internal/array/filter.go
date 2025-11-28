package array

// Filter iterates over a slice and calls a function (test) for each element.
// If test returns true, the element is added to the returned slice.
//
// The returned slice is a new slice, the original slice is not modified.
func Filter[T any](slice []T, test func(T) bool) (ret []T) {
	for _, item := range slice {
		if test(item) {
			ret = append(ret, item)
		}
	}

	return
}
