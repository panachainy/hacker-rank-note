package main

import (
	"fmt"
)

func process(arr []int32) int32 {
	var birds []int64 = make([]int64, 5)

	for _, v := range arr {
		switch v {
		case 1:
			birds[0]++
		case 2:
			birds[1]++
		case 3:
			birds[2]++
		case 4:
			birds[3]++
		case 5:
			birds[4]++
		}
	}
	maxIndex := 0
	var maxValue int32 = 0
	for i, v := range birds {
		// fmt.Println("i: ", i, "v: ", v)
		if maxValue < int32(v) {
			maxIndex = i
			maxValue = int32(v)
		}
	}

	return int32(maxIndex + 1)
}

func main() {
	var arr []int32 = []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4} // expect 3

	result := process(arr)
	fmt.Println(result)
}
