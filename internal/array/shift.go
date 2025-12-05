package array

// ShiftRet returns the first element of a slice and the rest of the slice
func ShiftRet[T any](slice []T) (item T, result []T) {
	item = slice[0]
	result = slice[1:]

	return
}

// Shift returns a new slice with the first element of the given slice removed
func Shift[T any](slice []T) (result []T) {
	_, result = ShiftRet(slice)

	return
}

// Unshift returns a new slice with the given item added to the front of the given slice
func Unshift[T any](slice []T, item T) (result []T) {
	result = append([]T{item}, slice...)

	return
}

// ShiftInPlace removes the first element of the given slice in place and returns it
func ShiftInPlace[T any](slice *[]T) (item T) {
	item = (*slice)[0]
	*slice = (*slice)[1:]

	return
}

// UnshiftInPlace adds the given item to the front of the given slice in place
func UnshiftInPlace[T any](slice *[]T, item T) {
	*slice = append([]T{item}, *slice...)
}
