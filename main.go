package main

import (
	"fmt"
)

func process(a []int32) int32 {
	var resultCount int32 = 0

	// # A
	// countHalf := len(a) / 2

	// firstArray := a[0:countHalf]
	// secondArray := a[countHalf:]

	// for _, firstValue := range firstArray {
	// 	for _, secondValue := range secondArray {
	// 		result := firstValue - secondValue
	// 		if result <= 1 {
	// 			fmt.Println("result ", firstValue, secondValue, result, true)
	// 			resultCount++
	// 		} else {
	// 			fmt.Println("result ", firstValue, secondValue, result, false)
	// 		}
	// 	}
	// }

	// # B
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if j == i {
				continue
			}
			result := a[i] - a[j]
			if result <= 1 && result >= 0 {
				fmt.Println("result ", a[i], a[j], result)
				resultCount++
			}
		}
	}

	// # C

	// var count int32 = 0
	// // max := 0
	// for i := 0; i < len(a); i++ {
	// 	for j := i; j < len(a); j++ {
	// 		if a[j]-a[i] <= 1 {
	// 			count++
	// 		}
	// 	}
	// 	if count > resultCount {
	// 		resultCount = count
	// 	}
	// 	count = 0
	// }

	return resultCount
}

func main() {
	// var a []int32 = []int32{1, 2, 2, 1, 2}
	// var a []int32 = []int32{1, 2, 2, 3, 1, 2}
	var a []int32 = []int32{4, 6, 5, 3, 3, 1} // expect 3

	result := process(a)
	fmt.Println(result)
}
