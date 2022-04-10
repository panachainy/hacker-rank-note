package main

import "fmt"

// https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true
func countingValleys(steps int32, path string) int32 {
	currentHigh := int32(0)
	valleyCount := int32(0)

	for _, v := range path {
		direction := string(v)

		if direction == "U" {
			currentHigh++
		} else {
			currentHigh--
		}

		if currentHigh == 0 && direction == "U" {
			valleyCount++
			fmt.Println("valleyCount ", valleyCount)
		}
	}

	return valleyCount
}

func main() {
	resp := countingValleys(100, "DDUDUDDUDDUDDUUUUDUDDDUUDDUUDDDUUDDUUUUUUDUDDDDUDDUUDUUDUDUUUDUUUUUDDUDDDDUDDUDDDDUUUUDUUDUUDUUDUDDD")
	fmt.Println("final expect:", 5, "result:", resp)
}
