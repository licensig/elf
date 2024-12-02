package elf

import (
	"strconv"
	"strings"
)

// StringAtoiArray splits a string to an array of int
func StringAtoiArray(str string, sep string) []int {
	// Create an array of string str by splitting it with sep
	strArr := strings.Split(str, sep)
	// Create an array of int with same length as strArr
	intArr := make([]int, len(strArr))

	// Convert each element of strArr to int and store it in intArr
	for i, s := range strArr {
		j, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intArr[i] = j
	}

	return intArr
}
