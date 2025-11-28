package array

// Map applies the given function fn to each element of the input slice
func Map[T any, U any](slice []T, fn func(idx int, element T) U) (ret []U) {
	for i, item := range slice {
		ret = append(ret, fn(i, item))
	}

	return
}
