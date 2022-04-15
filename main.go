package main

import (
	"fmt"
	"math"
	"sort"
)

func sortMin(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
}

func pickingNumbers(a []int32) int32 {
	var resultCount int32 = 0
	sortMin(a)

	for i := len(a) - 1; i >= 0; i-- {
		isCount := true

		//
		fmt.Println("start new loop ---")
		for j, v := range a {
			if i == j {
				fmt.Println("skip: ", i, j)
				continue
			}

			if math.Abs(float64(v-a[i])) > 1 {
				fmt.Println("not pass ", v, a[i], math.Abs(float64(v-a[i])))
				isCount = false
				break
			}
			fmt.Println("pass ", v, a[i], math.Abs(float64(v-a[i])))
		}

		//
		if isCount {
			fmt.Println("count: a = ", a[i])
			resultCount++
		}
	}

	// all element must <= 1

	return resultCount
}

func main() {
	fmt.Println("expect - 3", "result:",
		pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))

	fmt.Println("===========")

	fmt.Println("expect - 5", "result:",
		pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
}
