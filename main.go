package main

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/candies/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func minimumAbsoluteDifference(arr []int32) int32 {
	c := make(chan float64)
	go func(c chan<- float64, arr []int32) {
		minimum := math.MaxFloat64

		arr2 := arr

		for _, v := range arr {
			for _, v2 := range arr2 {
				if resAfterAb := math.Abs(float64(v - v2)); resAfterAb < minimum && resAfterAb > 0 {
					minimum = resAfterAb
				}
			}
		}

		c <- minimum

		close(c)
	}(c, arr)
	res := <-c

	return int32(res)
}

func main() {
	inputs := []int32{2222, 23123, 345566, 3312, -323236756}
	res := minimumAbsoluteDifference(inputs)
	fmt.Printf("done: %v\n", res)
}
