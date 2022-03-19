package main

func reverseArray(a []int32) []int32 {
	var b []int32

	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}

	return b
}
