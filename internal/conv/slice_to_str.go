package conv

import (
	"fmt"
	"strings"
)

// SliceToStr converts a slice to a string with a delimiter
func SliceToStr[T any](slice []T, delimiter string) string {
	return strings.Trim(
		strings.Join(
			strings.Fields(
				fmt.Sprintf("%v", slice),
			),
			delimiter,
		),
		"[]",
	)
}
