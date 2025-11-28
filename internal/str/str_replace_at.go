// Package str provides utility functions for string manipulation.
package str

// StrReplaceAt replaces the rune at index i in string s with rune r.
func StrReplaceAt(s string, i int, r rune) string {
	result := []rune(s)
	result[i] = r

	return string(result)
}
