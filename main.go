package main

import (
	"fmt"
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

func remove(s []int32, index int) []int32 {
	return append(s[:index], s[index+1:]...)
}

// func remove(s []int32, i int) []int32 {
// 	fmt.Println("============================")
// 	fmt.Println("1 remove s: ", s)
// 	s[i] = s[len(s)-1]
// 	fmt.Println("2 remove s: ", s)
// 	res := s[:len(s)-1]
// 	fmt.Println("3 remove s: ", s)
// 	return res
// }

// sort
// sort.Slice(number, func(i, j int) bool { return number[i] < number[j] })

// revert sort
// sort.Slice(number, func(i, j int) bool { return number[i] > number[j] })

func waiter(number []int32, q int32) []int32 {
	var res []int32
	// first loop
	// for i := 0; i < int(q); i++ {
	for i := 0; i < 9999; i++ {
		fmt.Println("==========================")

		if !checkPrimeNumber(i) {
			fmt.Println("skip")
			continue
		}

		// revert sort
		sort.Slice(number, func(i, j int) bool { return number[i] > number[j] })

		tmpNumber := number
		fmt.Println("i (prime): ", i)
		// fmt.Println("tmpN", tmpNumber)
		fmt.Println("N   ", number)

		// var tmpIndex []int

		// fmt.Println("number: ", i, number)
		// fmt.Println("tmpNumber: ", i, tmpNumber)

		// second loop from last to begin
		// for j, v := range tmpNumber {

		for j := len(tmpNumber) - 1; j >= 0; j-- {
			fmt.Println("in j loop ===", j)

			fmt.Println("tmpNumber[j] value: ", tmpNumber[j])
			fmt.Println("tmpNumber[j]%int32(i) value: ", tmpNumber[j]%int32(i))

			if tmpNumber[j]%int32(i) == 0 {
				fmt.Println("== in % loop ==")

				// TODO: add command
				res = append(res, tmpNumber[j])
				fmt.Println("% loop res: ", res)

				// TODO: remove command
				number = remove(number, j)
				fmt.Println("% loop num: ", number)
			}
		}

		fmt.Println("res", res)

		if len(number) == 0 {
			break
		}
	}

	// reverse before start check divisible by prime number
	// sort divisible value and move to answers

	return res
}

func main() {
	res := waiter([]int32{3, 4, 7, 6, 5}, 5)
	fmt.Println("res main: ", res)
}

// 4,6,3,7,5
