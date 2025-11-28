package str

// PadLeft pads the given string on the left with the specified pad string until it reaches the desired length.
func PadLeft(str string, length int, padStr string) string {
	for len(str) < length {
		str = padStr + str
	}

	return str
}

// PadRight pads the given string on the right with the specified pad string until it reaches the desired length.
func PadRight(str string, length int, padStr string) string {
	for len(str) < length {
		str = str + padStr
	}

	return str
}
