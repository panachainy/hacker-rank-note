package main

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/candies/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func minimumAbsoluteDifference(arr []int32) int32 {
	minimum := 9999999.9
	arr2 := arr

	for _, v := range arr {
		for _, v2 := range arr2 {
			if resAfterAb := math.Abs(float64(v - v2)); resAfterAb < minimum && resAfterAb > 0 {
				minimum = resAfterAb
			}
		}
	}

	return int32(minimum)
}

func main() {
	inputs := []int32{2222, 23123, 345566, 3312, -323236756}
	res := minimumAbsoluteDifference(inputs)
	fmt.Printf("done: %v\n", res)
}
