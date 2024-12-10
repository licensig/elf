package elf

import (
	"strconv"
	"strings"
)

// StringToInts converts a string of numbers separated by a delimiter into a slice of integers.
func StringToInts(str string, sep string) ([]int, error) {
	strArr := strings.Split(str, sep)
	intArr := make([]int, len(strArr))

	for i, s := range strArr {
		j, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		intArr[i] = j
	}

	return intArr, nil
}
