package conv

import "strconv"

// StrToInt converts a string representing a number to its integer value.
func StrToInt(numberString string) int {
	result, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}

	return result
}
