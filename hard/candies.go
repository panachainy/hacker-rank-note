// TODO: now it not pass this test
package main

import (
	"fmt"
)

// https://www.hackerrank.com/challenges/candies/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func candies(n int32, arr []int32) int64 {
	var candies []int32

	// Q:      2, 4, 2, 6, 1, 7, 8, 9, 2, 1

	// --:     1, 2, 1, 2, 1, 2, 3, 4, 1, 1
	// expect: 1, 2, 1, 2, 1, 2, 3, 4, 2, 1

	// sum:    19

	for i, value := range arr {
		if i == 0 {
			candies = append(candies, 1)
			continue
		}

		if value > arr[i-1] {
			candies = append(candies, candies[i-1]+1)
		} else {
			// check
			candies = append(candies, 1)
		}
	}

	// check value
	fmt.Printf("res v -> %v\n", candies)

	var sum int32

	for _, v := range candies {
		sum += v
	}

	return int64(sum)
}

func main() {
	inputs3 := []int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1}
	lenth3 := int32(len(inputs3))
	res3 := candies(lenth3, inputs3)

	fmt.Printf("res3 should 19: %v\n", res3)
}
