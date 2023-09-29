package main

import (
	"fmt"
)

func updateSliceAtIndex(slice []int32, index int32, newValue int32) []int32 {
	// Check if the index is within the bounds of the slice
	if index >= int32(0) && index < int32(len(slice)) {
		// Create a new slice by slicing the original slice
		// up to the specified index, then appending the new value,
		// and finally appending the rest of the original slice.
		newSlice := append(append([]int32{}, slice[:index]...), newValue)
		newSlice = append(newSlice, slice[index+1:]...)
		return newSlice
	} else {
		fmt.Println("Index out of range")
		return slice // Return the original slice if the index is out of range
	}
}

func permutationEquation(p []int32) []int32 {
	results := make([]int32, len(p))

	for _, item := range p {
		fmt.Println(item)
		var value = p[p[item-1]-1]
		fmt.Println("result: ", value, "index", value-1)

		results = updateSliceAtIndex(results, value-1, item)

		fmt.Println("done: ", results)
		fmt.Println("=====================")
	}

	return results
}

func main() {
	// fmt.Println("[[[res: ", permutationEquation([]int32{2, 3, 1}))
	fmt.Println("==")
	fmt.Println("[[[res: ", permutationEquation([]int32{4, 3, 5, 1, 2}))
	fmt.Println("==")
	// fmt.Println("[[[res: ", getTotalX(17, 24))
	// fmt.Println("==")
}
