package main

import (
	"strconv"
)

// https://www.hackerrank.com/challenges/find-digits/problem?isFullScreen=true
func findDigits(n int32) int32 {
	stringN := strconv.Itoa(int(n))

	resultCount := int32(0)

	for _, v := range stringN {
		vInt, _ := strconv.Atoi(string(v))

		if int32(vInt) == 0 {
			continue
		}
		// fmt.Println("before divine int32(vInt) ", int32(vInt))
		if n%int32(vInt) == 0 {
			resultCount++
		}
	}

	// fmt.Println("resultCount", resultCount)
	return resultCount
}
