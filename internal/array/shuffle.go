package array

import (
	"math/rand/v2"
	"slices"
)

// Shuffle returns a new slice with the elements of the input slice shuffled randomly.
func Shuffle[T any](s []T) []T {
	r := slices.Clone(s)

	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })

	return r
}
