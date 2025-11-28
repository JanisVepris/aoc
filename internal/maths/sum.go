package maths

// SumSliceInt takes a slice of integers and returns their sum.
func SumSliceInt(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
