package main

import "math"

func diagonalDifference(arr [][]int32) int32 {
	sum1 := int32(0)
	sum2 := int32(0)

	for i := 0; i < len(arr); i++ {
		sum1 += int32(arr[i][i])
		sum2 += int32(arr[0+i][(len(arr)-1)-i])
	}

	res := math.Abs(float64(sum1 - sum2))

	return int32(res)
}
