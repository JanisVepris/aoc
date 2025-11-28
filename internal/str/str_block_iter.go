package str

import "iter"

// StrBlockIterOverlap generates overlapping blocks of size rows x cols
// from the input slice of strings s. Each block is represented as a slice
// of strings, where each string corresponds to a row in the block.
// The function yields each block to the provided yield function.
// If the block size exceeds the dimensions of s, no blocks are yielded.
func StrBlockIterOverlap(s []string, rows, cols int) iter.Seq[[]string] {
	maxR := len(s)

	maxC := 0
	if maxR > 0 {
		maxC = len(s[0])
	}

	maxRowBound := maxR - rows + 1
	maxColBound := maxC - cols + 1

	return func(yield func([]string) bool) {
		if rows > maxR || cols > maxC || rows < 1 || cols < 1 {
			return
		}

		block := make([]string, rows)

		for r := range maxRowBound {
			for c := range maxColBound {
				for i := range rows {
					block[i] = s[r+i][c : c+cols]
				}

				if !yield(block) {
					return
				}
			}
		}
	}
}
