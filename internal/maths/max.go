package maths

// MaxInt returns the maximum integer from a variable number of integer arguments.
func MaxInt(nums ...int) int {
	if len(nums) == 0 {
		panic("MaxInts requires at least one argument")
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	return max
}
