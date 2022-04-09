package main

import (
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/maximum-element/problem?isFullScreen=true
func getMaxFrom(arrays []int32) int32 {
	max := int32(0)
	for _, v := range arrays {
		if max < v {
			max = v
		}
	}

	return max
}

func remove(slice []int32, s int) []int32 {
	return append(slice[:s], slice[s+1:]...)
}

func getMax(operations []string) []int32 {
	stacks := []int32{}
	result := []int32{}
	for _, v := range operations {
		switch firstChar := v[:1]; firstChar {
		// Push the element x into the stack.
		case "1":
			spliteds := strings.Split(v, " ")
			secondValue, _ := strconv.Atoi(spliteds[1])

			stacks = append(stacks, int32(secondValue))

		// delete top of stack
		case "2":
			stacks = remove(stacks, len(stacks)-1)

		// Print the maximum element in the stack.
		case "3":
			result = append(result, getMaxFrom(stacks))
		}
	}

	return result
}
