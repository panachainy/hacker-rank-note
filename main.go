package main

import (
	"fmt"
)

func checkArea(s int32, t int32, treePoint int32, fruits []int32) int32 {
	var count int32 = 0

	for _, fruit := range fruits {

		realFruitPoint := treePoint + fruit

		// Sam house
		if realFruitPoint >= s && realFruitPoint <= t {
			count++
		}
	}
	return count
}

func process(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	appleTree := a
	orageTree := b

	appleInSamHouse := checkArea(s, t, appleTree, apples)
	orageInSamHouse := checkArea(s, t, orageTree, oranges)

	fmt.Println(appleInSamHouse)

	fmt.Println(orageInSamHouse)
}

func main() {
	var s int32 = 7
	var t int32 = 11
	var a int32 = 5
	var b int32 = 15
	var apples []int32 = []int32{-2, 2, 1}
	var oranges []int32 = []int32{5, -6}

	process(s, t, a, b, apples, oranges)
}
