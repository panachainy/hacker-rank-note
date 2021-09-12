package main

import (
	"fmt"
)

func absDiffInt(x, y int32) int32 {
	if x < y {
		return y - x
	}
	return x - y
}

func process(x int32, y int32, z int32) string {
	var x2z int32 = absDiffInt(x, z)
	var y2z int32 = absDiffInt(y, z)

	if x2z < y2z {
		return "Cat A"
	} else if x2z > y2z {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {
	var x int32 = 1
	var y int32 = 2
	var z int32 = 3

	result := process(x, y, z)
	fmt.Println(result)
}
