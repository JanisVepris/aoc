package day10

import (
	"fmt"
	"math"
	"math/bits"
	"regexp"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

type Machine struct {
	lights         int      // n of lights
	target         uint64   // target light pattern as bitmask
	buttons        []uint64 // button[j] = bitmask of lights toggled by button j (for Part 1)
	joltages       []int    // target joltage levels
	joltageButtons [][]int  // for Part 2: buttons[b] -> which joltages this button increments
}

var lines []string

func Setup() {
	lines = files.ReadFile("2025/day10/input.txt")
}

func Part1() {
	result := 0
	machines := parseMachines()

	for _, machine := range machines {
		_, count, _ := CalcLights(machine.buttons, machine.target, machine.lights)
		result += count
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	machines := parseMachines()

	for _, machine := range machines {
		presses, ok := CalcJoltages(machine.joltages, machine.buttons)
		if !ok {
			continue
		}
		result += presses
	}

	fmt.Printf("Part 2: %d\n", result)
}

func parseMachines() []Machine {
	lightsRegex := regexp.MustCompile(`\[([#.]*)\]`)
	buttonsRegex := regexp.MustCompile(`\(([\d,]*)\)`)
	joltageRegex := regexp.MustCompile(`{([\d,]*)}`)

	machines := make([]Machine, len(lines))

	for m, line := range lines {
		// target pattern
		lightsStr := lightsRegex.FindStringSubmatch(line)[1]

		var targetMask uint64
		for i, ch := range lightsStr {
			if ch == '#' {
				targetMask |= 1 << i
			}
		}
		lightCount := len(lightsStr)

		btnMatches := buttonsRegex.FindAllStringSubmatch(line, -1)
		buttonMasks := make([]uint64, len(btnMatches))
		joltageButtons := make([][]int, len(btnMatches))

		for i, match := range btnMatches {
			lightNums := strings.Split(match[1], ",")
			var mask uint64

			for _, lightNum := range lightNums {
				if lightNum == "" {
					continue
				}

				idx := conv.StrToInt(lightNum)
				mask |= 1 << idx

				joltageButtons[i] = append(joltageButtons[i], idx)
			}

			buttonMasks[i] = mask
		}

		joltageStr := joltageRegex.FindStringSubmatch(line)[1]
		joltageNums := strings.Split(joltageStr, ",")

		joltages := make([]int, 0, len(joltageNums))
		for _, joltageNum := range joltageNums {
			if joltageNum == "" {
				continue
			}
			joltages = append(joltages, conv.StrToInt(joltageNum))
		}

		machines[m] = Machine{
			lights:         lightCount,
			target:         targetMask,
			buttons:        buttonMasks,
			joltages:       joltages,
			joltageButtons: joltageButtons,
		}
	}

	return machines
}

// P1: lights (GF(2) solver)
func CalcLights(buttons []uint64, target uint64, lights int) (bestMask uint64, bestCount int, ok bool) {
	m := len(buttons)

	// rows[r] encodes equation for light r:
	//   - lower m bits: coefficients for each button
	//   - bit m: RHS (target[r])
	rows := make([]uint64, lights)
	for light := range lights {
		var row uint64
		for btn := range m {
			if (buttons[btn]>>light)&1 == 1 {
				row |= 1 << btn
			}
		}
		if (target>>light)&1 == 1 {
			row |= 1 << m // RHS bit
		}
		rows[light] = row
	}

	// Gaussian elimination to RREF over GF(2)
	pivotOfCol := make([]int, m)
	for i := range pivotOfCol {
		pivotOfCol[i] = -1
	}

	row := 0
	for col := 0; col < m && row < lights; col++ {
		// find pivot row with 1 in this column at or below 'row'
		pivot := -1
		for r := row; r < lights; r++ {
			if (rows[r]>>col)&1 == 1 {
				pivot = r
				break
			}
		}
		if pivot == -1 {
			// free column (no pivot)
			continue
		}

		// swap current row with pivot row
		rows[row], rows[pivot] = rows[pivot], rows[row]
		pivotOfCol[col] = row

		// eliminate this column in all other rows
		for r := range lights {
			if r != row && ((rows[r]>>col)&1 == 1) {
				rows[r] ^= rows[row]
			}
		}

		row++
	}

	// collect free columns (no pivot)
	var freeCols []int
	for col := range m {
		if pivotOfCol[col] == -1 {
			freeCols = append(freeCols, col)
		}
	}
	k := len(freeCols) // nullspace dimension

	// particular solution x0: set all free vars = 0 and read pivots from RHS
	var x0 uint64
	for col := range m {
		rowIdx := pivotOfCol[col]
		if rowIdx == -1 {
			continue // free var, stays 0
		}
		rhs := (rows[rowIdx] >> m) & 1
		if rhs == 1 {
			x0 |= 1 << col
		}
	}

	// build nullspace basis: one vector per free variable
	basis := make([]uint64, 0, k)
	for _, fcol := range freeCols {
		var v uint64
		// this free var set to 1
		v |= 1 << fcol

		// pivot vars depend on it via their row equations
		for col := range m {
			rowIdx := pivotOfCol[col]
			if rowIdx == -1 {
				continue
			}
			coeff := (rows[rowIdx] >> fcol) & 1
			if coeff == 1 {
				v |= 1 << col
			}
		}

		basis = append(basis, v)
	}

	// search over all combinations x = x0 XOR sum(c_i * basis[i])
	limit := uint64(1)
	if k > 0 {
		limit = 1 << k
	}

	bestMask = 0
	bestCount = 0
	ok = false

	for mask := uint64(0); mask < limit; mask++ {
		x := x0
		for i := range k {
			if (mask>>i)&1 == 1 {
				x ^= basis[i]
			}
		}
		pop := bits.OnesCount64(x)
		if !ok || pop < bestCount {
			ok = true
			bestCount = pop
			bestMask = x
		}
	}

	return bestMask, bestCount, ok
}

// helper: check A * x == target exactly (integer arithmetic)
func checkSolutionInt(A [][]int, target []int, x []int) bool {
	k := len(target)
	m := len(x)
	for i := range k {
		sum := 0
		row := A[i]
		for j := range m {
			sum += row[j] * x[j]
		}
		if sum != target[i] {
			return false
		}
	}
	return true
}

//  1. Build A (k x m) with 0/1 entries and RHS = target.
//  2. Do Gaussian elimination over float64 to find pivot vs free variables.
//  3. There are at most a few free variables
//  4. Brute-force those free vars in [0 .. per-var max] and solve pivots
//     by back-substitution; keep the integer, non-negative solution with minimal sum.
func CalcJoltages(target []int, buttonMasks []uint64) (int, bool) {
	k := len(target)
	m := len(buttonMasks)

	if k == 0 {
		return 0, true
	}
	if m == 0 {
		for _, t := range target {
			if t != 0 {
				return 0, false
			}
		}
		return 0, true
	}

	// Build integer A and float augmented matrix [A | target]
	Aint := make([][]int, k)
	A := make([][]float64, k)
	for i := range k {
		Aint[i] = make([]int, m)
		A[i] = make([]float64, m+1)
		for j := range m {
			val := 0
			if (buttonMasks[j]>>i)&1 == 1 {
				val = 1
			}
			Aint[i][j] = val
			A[i][j] = float64(val)
		}
		A[i][m] = float64(target[i])
	}

	// Gaussian elimination to row-echelon form over R
	pivotColForRow := make([]int, k)
	for i := range pivotColForRow {
		pivotColForRow[i] = -1
	}

	row := 0
	const eps = 1e-9

	for col := 0; col < m && row < k; col++ {
		// partial pivoting
		best := -1
		bestAbs := 0.0
		for r := row; r < k; r++ {
			val := math.Abs(A[r][col])
			if val > eps && val > bestAbs {
				bestAbs = val
				best = r
			}
		}
		if best == -1 {
			// no pivot in this column
			continue
		}

		// swap into current row
		A[row], A[best] = A[best], A[row]

		// normalize pivot row so pivot = 1
		pivotVal := A[row][col]
		inv := 1.0 / pivotVal
		for c := col; c <= m; c++ {
			A[row][c] *= inv
		}

		// eliminate rows BELOW
		for r := row + 1; r < k; r++ {
			factor := A[r][col]
			if math.Abs(factor) <= eps {
				continue
			}
			for c := col; c <= m; c++ {
				A[r][c] -= factor * A[row][c]
			}
		}

		pivotColForRow[row] = col
		row++
	}

	rank := row

	// Identify pivot and free columns
	isPivotCol := make([]bool, m)
	for r := range rank {
		c := pivotColForRow[r]
		if c >= 0 {
			isPivotCol[c] = true
		}
	}

	var freeCols []int
	for j := range m {
		if !isPivotCol[j] {
			freeCols = append(freeCols, j)
		}
	}
	kfree := len(freeCols)

	// Bounds for each free variable:
	// x_j can't exceed any target that it contributes to.
	freeMax := make([]int, kfree)
	for i := range freeMax {
		freeMax[i] = math.MaxInt32
	}
	for fi, col := range freeCols {
		maxVal := math.MaxInt32
		for rowIdx := range k {
			if Aint[rowIdx][col] == 1 {
				if target[rowIdx] < maxVal {
					maxVal = target[rowIdx]
				}
			}
		}
		if maxVal == math.MaxInt32 {
			// button doesn't affect any joltage: never worth pressing
			maxVal = 0
		}
		freeMax[fi] = maxVal
	}

	// Quick path: no free variables => unique solution
	if kfree == 0 {
		x := make([]float64, m)
		for r := rank - 1; r >= 0; r-- {
			col := pivotColForRow[r]
			if col < 0 {
				continue
			}
			sum := 0.0
			for c := col + 1; c < m; c++ {
				sum += A[r][c] * x[c]
			}
			x[col] = A[r][m] - sum
		}

		xi := make([]int, m)
		total := 0
		for j := range m {
			v := x[j]
			rnd := math.Round(v)
			if math.Abs(v-rnd) > 1e-6 || rnd < 0 {
				return 0, false
			}
			xi[j] = int(rnd)
			total += xi[j]
		}

		if !checkSolutionInt(Aint, target, xi) {
			return 0, false
		}
		return total, true
	}

	// General case: some free variables
	bestSum := math.MaxInt32
	bestFound := false

	freeVals := make([]int, kfree)
	x := make([]float64, m)
	xi := make([]int, m)

	var dfs func(idx int, partialFreeSum int)
	dfs = func(idx int, partialFreeSum int) {
		if bestFound && partialFreeSum >= bestSum {
			// even if all pivots were 0, we can't beat current best
			return
		}

		if idx == kfree {
			// set free vars in x
			for j := range m {
				x[j] = 0
			}
			for fi, col := range freeCols {
				x[col] = float64(freeVals[fi])
			}

			// back-substitute to compute pivot variables
			for r := rank - 1; r >= 0; r-- {
				col := pivotColForRow[r]
				if col < 0 {
					continue
				}
				sum := 0.0
				for c := col + 1; c < m; c++ {
					sum += A[r][c] * x[c]
				}
				x[col] = A[r][m] - sum
			}

			total := partialFreeSum
			for j := range m {
				v := x[j]
				rnd := math.Round(v)
				if math.Abs(v-rnd) > 1e-6 || rnd < 0 {
					return
				}
				xi[j] = int(rnd)
				if j >= 0 && !isPivotCol[j] {
					// free vars already counted in partialFreeSum
					continue
				}
				total += xi[j]
				if bestFound && total >= bestSum {
					return
				}
			}

			if !checkSolutionInt(Aint, target, xi) {
				return
			}

			if !bestFound || total < bestSum {
				bestFound = true
				bestSum = total
			}
			return
		}

		maxVal := freeMax[idx]
		for v := 0; v <= maxVal; v++ {
			newSum := partialFreeSum + v
			if bestFound && newSum >= bestSum {
				break // higher v will only increase sum
			}
			freeVals[idx] = v
			dfs(idx+1, newSum)
		}
	}

	dfs(0, 0)

	if !bestFound {
		return 0, false
	}
	return bestSum, true
}
