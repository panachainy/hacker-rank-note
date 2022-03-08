package main

import "fmt"

func candies(n int32, arr []int32) int64 {
	var candies []int32

	for i, value := range arr {
		if i == 0 {
			candies = append(candies, 1)
			continue
		}

		if value > arr[i-1] {
			candies = append(candies, candies[i-1]+1)
		} else {
			candies = append(candies, 1)
		}
	}

	// check value
	for _, v := range candies {
		fmt.Printf("res v -> %v\n", v)
	}

	// 4, 6, 4, 5, 6, 2
	// 1, 2, 1, 2, 3, 1
	var sum int32

	for _, v := range candies {
		sum += v
	}

	return int64(sum)
}

func main() {
	// inputs := []int32{4, 6, 4, 5, 6, 2}
	// lenth := int32(len(inputs))
	// res1 := candies(lenth, inputs)

	// fmt.Printf("res1 should 10: %v\n", res1)

	// inputs2 := []int32{3, 1, 2, 2}
	// lenth2 := int32(len(inputs2))
	// res2 := candies(lenth2, inputs2)

	// fmt.Printf("res2 should 4: %v\n", res2)

	// inputs3 := []int32{10, 2, 4, 2, 6, 1, 7, 8, 9, 2, 1}
	// lenth3 := int32(len(inputs3))
	// res3 := candies(lenth3, inputs3)

	// fmt.Printf("res3 should 19: %v\n", res3)

	inputs4 := []int32{8, 2, 4, 3, 5, 2, 6, 4, 5}
	lenth4 := int32(len(inputs4))
	res4 := candies(lenth4, inputs4)

	fmt.Printf("res4 should 12: %v\n", res4)

	// 12121212
}
