package main

import (
	"fmt"
	"strconv"
)

func process(p int32, q int32) {
	result := []int32{}

	for i := p; i <= q; i++ {
		p2 := i * i
		p2String := strconv.Itoa(int(p2))
		p2Length := len(p2String)

		halfP2 := p2Length / 2

		p2First := p2String[:halfP2]
		p2Second := p2String[halfP2:p2Length]

		// string to int
		p2FirstInt, _ := strconv.Atoi(p2First)
		p2SecondInt, _ := strconv.Atoi(p2Second)

		temp := int32(p2FirstInt) + int32(p2SecondInt)

		if temp == i {
			result = append(result, i)
		}

	}

	if len(result) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		for i, v := range result {
			fmt.Print(v)
			if i != len(result)-1 {
				fmt.Print(" ")
			}
		}
	}
}

func main() {
	var p int32 = 1
	var q int32 = 100

	process(p, q)
}
