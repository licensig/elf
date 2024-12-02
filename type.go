package elf

import (
	"strconv"
	"strings"
)

// StringToInts converts a string of numbers separated by a delimiter into a slice of integers.
// Panics if any substring cannot be parsed as a valid integer.
func StringToInts(str string, sep string) []int {
	strArr := strings.Split(str, sep)
	intArr := make([]int, len(strArr))

	for i, s := range strArr {
		j, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intArr[i] = j
	}

	return intArr
}
