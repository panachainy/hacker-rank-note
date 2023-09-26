package main

import (
	"fmt"
)

func utopianTree(n int32) int32 {
	height := int32(0)
	isDouble := false
	fmt.Println("n", n)

	for i := int32(0); i <= n; i++ {
		fmt.Println("i", i)

		if isDouble {
			height *= 2
			fmt.Println("x2")
		} else {
			height += 1
			fmt.Println("+1")

		}
		fmt.Println("height", height)

		isDouble = !isDouble
	}

	return height
}

func main() {
	fmt.Println("[[[res: ", utopianTree(0))
	fmt.Println("==")
	fmt.Println("[[[res: ", utopianTree(1))
	fmt.Println("==")
	fmt.Println("[[[res: ", utopianTree(3))
	fmt.Println("==")
	fmt.Println("[[[res: ", utopianTree(4))
}
