package main

import (
	"fmt"
	"sort"
)

func sortMoreToLess(datas []int32) []int32 {
	sort.Slice(datas, func(i, j int) bool { return datas[i] < datas[j] })
	return datas
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	max := int32(0)

	keyboards = sortMoreToLess(keyboards)
	drives = sortMoreToLess(drives)

	fmt.Println("b", b)

	for _, vk := range keyboards {
		for _, vd := range drives {
			tmpSum := vk + vd
			fmt.Println("test ====", tmpSum)

			if tmpSum <= b && tmpSum > max {
				fmt.Println("in if ====", tmpSum, b, max)
				max = tmpSum
			}
		}
	}

	if max == 0 {
		max = -1
	}

	return max
}
