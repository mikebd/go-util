package shell

import (
	"strconv"
)

// IsOutputGreaterThanZero returns true if the null terminated output is a number greater than zero
func IsOutputGreaterThanZero(output []byte) bool {
	if len(output) <= 1 {
		return false
	}
	value, err := strconv.Atoi(string(output)[:len(output)-1])
	if err != nil {

	}
	return value > 0
}
