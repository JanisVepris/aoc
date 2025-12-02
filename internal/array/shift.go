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

func Unshift[T any](slice []T, item T) (result []T) {
	result = append([]T{item}, slice...)

	return
}
