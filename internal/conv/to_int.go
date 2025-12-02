package conv

import (
	"fmt"
	"strconv"
)

func ToInt(in any) int {
	switch v := in.(type) {
	case int:
		return v
	case string:
		val, _ := strconv.Atoi(v)
		return val
	case rune:
		return int(v - '0')
	case uint8:
		return int(v - '0')
	default:
		panic(fmt.Sprintf("cannot convert %T to int", v))
	}
}
