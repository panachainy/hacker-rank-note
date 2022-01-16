package main

import (
	"fmt"
)

func stringReduction(s string) int {
	a := 0
	b := 0
	c := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			a++
		case 'b':
			b++
		case 'c':
			c++
		}
	}

	fmt.Println(s)
	fmt.Println(a, b, c)

	// check empty more than 2 chars
	if (a == 0 && b == 0) || (b == 0 && c == 0) || (a == 0 && c == 0) {
		return len(s)
	}

	return len(s)
}

func main() {
	var strings []string = []string{"cab", "bcab", "ccccc"}
	var strings2 []string = []string{"abcbcba", "ababbac", "abababaccc"}

	for _, s := range strings {
		stringReduction(s)
		// fmt.Println(stringReduction(s))
	}
	// 2
	// 1
	// 5

	fmt.Println("==========================")

	for _, s := range strings2 {
		stringReduction(s)
		// fmt.Println(stringReduction(s))
	}
	// 1
	// 2
	// 1
}
