package main

import (
	"fmt"
	"strconv"
)

func gradeRounding(grade int32) int32 {
	if grade < 38 || grade%5 < 3 {
		return grade
	} else {
		return grade + 5 - grade%5
	}
}

func process(grades []int32) []int32 {
	var gradesRounded []int32

	for _, grade := range grades {
		gradesRounded = append(gradesRounded, gradeRounding(grade))
	}

	return gradesRounded
}

func main() {
	var grades []int32 = []int32{4, 73, 67, 38, 33}
	result := process(grades)

	fmt.Println("result")
	fmt.Println(result)
}

// convert string to int
func toInt(s string) int {
	intValue, _ := strconv.Atoi(s)
	return intValue
}

// // convert int to string
// func toString(i int) string {
// 	return strconv.Itoa(i)
// }

// convert int32 to string
func toString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}
