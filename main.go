package main

import (
	"fmt"
	"strings"
)

// func stringReduction(s string) int {
// 	a := 0
// 	b := 0
// 	c := 0

// 	for i := 0; i < len(s); i++ {
// 		switch firstValue {
// 		case 'a':
// 			a++
// 		case 'b':
// 			b++
// 		case 'c':
// 			c++
// 		}
// 	}

// 	fmt.Println(s)
// 	fmt.Println(a, b, c)

// 	// check empty more than 2 chars
// 	if (a == 0 && b == 0) || (b == 0 && c == 0) || (a == 0 && c == 0) {
// 		return len(s)
// 	}

// 	return len(s)
// }

var (
	A = "a"
	B = "b"
	C = "c"
)

func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:]) // Shift a[i+1:] left one index.
	s[len(s)-1] = ""     // Erase last element (write zero value).
	s = s[:len(s)-1]     // Truncate slice.

	return s
}

// 0 <= index <= len(a)
func insert(s []string, index int, value string) []string {
	if len(s) == index { // nil or empty slice or after last element
		return append(s, value)
	}
	s = append(s[:index+1], s[index:]...) // index < len(a)
	s[index] = value
	return s
}

func stringReductionProcess(s string) string {
	chars := strings.Split(s, "")

	if len(chars) == 1 {
		return s
	}

	if len(chars) == 2 && chars[0] == chars[1] {
		return s
	}

	for i := 0; i < len(chars)-1; i++ {
		newChar := ""

		firstValue := string(chars[i])
		secondValue := string(chars[i+1])

		fmt.Println(string(firstValue), string(secondValue))

		if firstValue == secondValue {
			continue
		}

		// Get third value
		if (firstValue == A && secondValue == B) || (firstValue == B && secondValue == A) {
			newChar = C
		} else if (firstValue == A && secondValue == C) || (firstValue == C && secondValue == A) {
			newChar = B
		} else {
			newChar = A
		}

		fmt.Println("raw", chars)

		chars = remove(chars, i)
		fmt.Printf("1 %#v \n", chars)
		chars = remove(chars, i)
		fmt.Printf("2 %#v \n", chars)

		chars = insert(chars, i, newChar)
		fmt.Printf("3 %#v \n", chars)

		fmt.Println("new chars", newChar)
		fmt.Println("==========================")
	}

	if len(s) != len(chars) {
		return stringReductionProcess(strings.Join(chars, ""))
	}

	return strings.Join(chars, "")
}

func stringReduction(s string) int {
	reduced := stringReductionProcess(s)

	return len(reduced)
}

func main() {
	var strings []string = []string{"cab", "bcab", "ccccc", "abcbcba", "ababbac", "abababaccc"}

	// var strings []string = []string{"abababaccc"}

	for _, s := range strings {
		fmt.Println("Length ========", stringReduction(s))
	}
	// 2
	// 1
	// 5

	// 1
	// 2
	// 1
}
