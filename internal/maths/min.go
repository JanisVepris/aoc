package maths

// MinInt returns the smallest of the provided integers.
func MinInt(nums ...int) int {
	if len(nums) == 0 {
		panic("MinInt requires at least one argument")
	}

	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min
}
