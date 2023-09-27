package main

import (
	"fmt"
)

func selfDouble(number int32) int32 {
	return number * number
}

func squares(a int32, b int32) int32 {
	var squareIndex = int32(1)
	var squareCount = int32(0)

	for true {
		var currentDouble = selfDouble(squareIndex)

		if currentDouble >= a && currentDouble <= b {
			squareCount++
		}

		if currentDouble >= b {
			break
		}

		squareIndex++
	}

	return squareCount
}

func main() {
	fmt.Println("[[[res: ", squares(3, 9))
	fmt.Println("==")
	fmt.Println("[[[res: ", squares(17, 24))
	fmt.Println("==")
}
