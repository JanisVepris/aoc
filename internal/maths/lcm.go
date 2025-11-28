package maths

// LCM returns the least common multiple of a set of integers
func LCM(numbers ...int) int {
	if len(numbers) == 0 {
		return 1
	}

	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		gcdResult := GCD(result, numbers[i])
		result = result * numbers[i] / gcdResult
	}

	return result
}
