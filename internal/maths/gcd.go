package maths

// GCD returns the greatest common divisor of a(int) and b(int)
func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}
