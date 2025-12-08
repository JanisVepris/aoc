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

// StrToFloat64 converts a string representing a number to its float64 value.
func StrToFloat64(numberString string) float64 {
	result, err := strconv.ParseFloat(numberString, 64)
	if err != nil {
		panic(err)
	}

	return result
}
