package main

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/2d-array/problem?isFullScreen=true
func hourglassSum(arr [][]int32) int32 {
	hightest := int32(math.MinInt32)

	for i, v := range arr {
		if i > 3 {
			break
		}

		for j := range v {
			if j > 3 {
				break
			}

			sum := getHourGlassSum(arr, i, j)

			if sum > hightest {
				hightest = sum
			}
		}
	}

	return hightest
}

func getHourGlassSum(arr [][]int32, startI int, startJ int) int32 {
	sum := int32(0)

	for i := startI; i < startI+3; i++ {
		for j := startJ; j < startJ+3; j++ {
			if i == startI+1 {
				if j == startJ || j == startJ+2 {
					continue
				}
			}

			sum += arr[i][j]
		}
	}

	return sum
}

func main() {
	// expect 19

	inputs := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	res := hourglassSum(inputs)

	fmt.Printf("done: %v\n", res)
}
