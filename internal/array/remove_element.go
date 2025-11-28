package array

// RemoveElement removes the element at the specified index from the slice
func RemoveElement[T any](slice []T, index int) []T {
	copySlice := append([]T(nil), slice...)

	return append(copySlice[:index], copySlice[index+1:]...)
}
