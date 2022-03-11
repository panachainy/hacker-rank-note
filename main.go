package main

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/candies/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func minimumAbsoluteDifference(arr []int32) int32 {
	minimum := 9999999.9
	arr2 := arr

	c := make(chan float64)

	// count := 0

	for i, v := range arr {
		// count++
		go func(i int, v int32, c chan float64) {
			inMinimum := 9999999.9
			for i2, v2 := range arr2 {
				if i == i2 {
					continue
				}

				if resAfterAb := math.Abs(float64(v - v2)); resAfterAb < inMinimum && resAfterAb > 0 {
					inMinimum = resAfterAb
				}
			}
			c <- inMinimum
		}(i, v, c)
	}

	// result := <-c
	// fmt.Printf("result xxxxxx: %v\n", result)
	// fmt.Printf("count xxxxxx: %v\n", count)

	// for i := 0; i <= count; i++ {
	// 	fmt.Printf("i xxxxxx: %v\n", i)

	// 	result := <-c
	// 	fmt.Printf("result xxxxxx: %v\n", result)

	// 	// if result := <-c; result < minimum {
	// 	// 	minimum = result
	// 	// }
	// }

	for v := range c {
		fmt.Printf("received: %v\n", v)
	}

	fmt.Printf("done ===========\n")

	// return int32(math.Round(res))
	return int32(minimum)
}

func main() {
	inputs := []int32{2222, 23123, 345566, 3312, -323236756}
	res := minimumAbsoluteDifference(inputs)
	fmt.Printf("done: %v\n", res)
}
