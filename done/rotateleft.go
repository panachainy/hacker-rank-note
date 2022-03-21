package main

// https://www.hackerrank.com/challenges/array-left-rotation/problem?isFullScreen=true
// this is main func
func rotateLeft(d int32, arr []int32) []int32 {
	for i := 0; i < int(d); i++ {
		removedValue := arr[0]

		arr = remove(arr, 0)

		arr = append(arr, removedValue)
	}

	return arr
}

func remove(slice []int32, s int) []int32 {
	return append(slice[:s], slice[s+1:]...)
}
