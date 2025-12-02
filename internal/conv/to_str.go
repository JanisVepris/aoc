package conv

import (
	"fmt"
	"strconv"
)

// ToStr converts various types to their string representation.
func ToStr(in any) string {
	switch v := in.(type) {
	case string:
		return v
	case rune:
		return string(v)
	case byte:
		return string(rune(v))
	case int, int8, int16, int64:
		return strconv.Itoa(v.(int))
	case uint, uint16, uint32, uint64:
		return strconv.Itoa(v.(int))
	default:
		panic(fmt.Sprintf("cannot convert %T to string", v))
	}
}
