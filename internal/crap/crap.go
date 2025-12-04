package crap

import (
	"fmt"
	"time"
)

func PrintDuration(duration time.Duration) {
	us := duration.Microseconds()

	minutes := us / 60000000
	us %= 60000000

	seconds := us / 1000000
	us %= 1000000

	milliseconds := us / 1000
	us %= 1000

	fmt.Printf("(in %dm %ds %dms %dµs)\n", minutes, seconds, milliseconds, us)
}
