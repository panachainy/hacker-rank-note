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

		fmt.Println("p2Length: ", p2Length)
		fmt.Println("p2Length%2: ", p2Length%2)

		if p2Length%2 != 0 && p2Length > 1 {
			continue
		}

		p2First := p2String[:p2Length/2]
		p2Second := p2String[p2Length/2 : p2Length]

		fmt.Println("p2First: ", p2First)
		fmt.Println("p2Second: ", p2Second)

		// string to int32
		p2FirstInt, _ := strconv.Atoi(p2First)
		p2SecondInt, _ := strconv.Atoi(p2Second)

		temp := int32(p2FirstInt) + int32(p2SecondInt)

		// p2Array := []rune(p2String)

		// var temp int32 = 0

		// for _, v := range p2Array {
		// 	// rune to int32
		// 	valueInt, _ := strconv.Atoi(string(v))

		// 	temp += int32(valueInt)
		// }

		if temp == i {
			fmt.Println("Kaprekar")

			fmt.Println(i, temp)

			result = append(result, i)
		} else {
			fmt.Println("not Kaprekar")
			fmt.Println(i, temp)
		}
		fmt.Println("==========")
	}

	if len(result) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		fmt.Println(result)
	}
}

func main() {
	var p int32 = 1
	var q int32 = 100

	process(p, q)
}
