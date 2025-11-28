package maths

// MultiplySliceInt takes a slice of integers and returns their product.
func MultiplySliceInt(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result *= num
	}
	return result
}
