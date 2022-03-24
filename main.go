package main

import (
	"math"
	"sort"
)

func checkPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func remove(s []int32, i int) []int32 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// sort
// sort.Slice(number, func(i, j int) bool { return number[i] < number[j] })

// revert sort
// sort.Slice(number, func(i, j int) bool { return number[i] > number[j] })

func waiter(number []int32, q int32) []int32 {
	// first loop
	for i := 0; i < int(q); i++ {
		if checkPrimeNumber(i) {
			continue
		}

		// revert sort
		sort.Slice(number, func(i, j int) bool { return number[i] > number[j] })

		tmpNumber := number

		// second loop
		for j, v := range tmpNumber {
			if v%int32(i) == 0 {
				remove(number, j)
			}
		}

		if len(number) == 0 {
			break
		}
	}

	// reverse before start check divisible by prime number
	// sort divisible value and move to answers

	return number
}

func main() {
	waiter([]int32{3, 4, 7, 6, 5}, 5)
}

// 4,6,3,7,5
