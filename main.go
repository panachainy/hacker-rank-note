package main

import (
	"fmt"
	"math"
	"sort"
)

// https://www.hackerrank.com/challenges/candies/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func minimumAbsoluteDifference(arr []int32) int32 {
	minimum := math.MaxFloat64
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	lenth := len(arr)

	for i := range arr {
		if i == lenth-1 {
			break
		}

		if resAfterAb := math.Abs(float64(arr[i] - arr[i+1])); resAfterAb < minimum {
			minimum = resAfterAb
		}
	}

	return int32(minimum)
}

func main() {
	inputs := []int32{2222, 23123, 345566, 3312, -323236756}
	res := minimumAbsoluteDifference(inputs)
	fmt.Printf("done: %v\n", res)
}
