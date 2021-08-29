package main

import "fmt"

func process(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	}

	for x1 <= x2 {
		x1 += v1
		x2 += v2

		if x1 == x2 {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	var x1 int32 = 0
	var v1 int32 = 3
	var x2 int32 = 4
	var v2 int32 = 2

	result := process(x1, v1, x2, v2)
	fmt.Println(result)
}
