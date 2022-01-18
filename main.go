package main

import (
	"fmt"
	"strings"
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

func getThirdChar(firstValue string, secondValue string) string {
	A := "a"
	B := "b"
	C := "c"

	if (firstValue == A && secondValue == B) || (firstValue == B && secondValue == A) {
		return C
	} else if (firstValue == A && secondValue == C) || (firstValue == C && secondValue == A) {
		return B
	} else {
		return A
	}
}

func stringReductionProcess(s string, startIndex int) string {
	chars := strings.Split(s, "")

	if len(chars) == 1 {
		return s
	}

	for i := startIndex; i < len(chars)-1; i++ {
		firstValue := string(chars[i])
		secondValue := string(chars[i+1])

		if firstValue == secondValue {
			continue
		}

		newChar := getThirdChar(firstValue, secondValue)

		chars = remove(chars, i)
		chars = remove(chars, i)
		chars = insert(chars, i, newChar)
	}

	if len(s) != len(chars) {
		return stringReductionProcess(strings.Join(chars, ""), 0)
	}

	return strings.Join(chars, "")
}

func stringReduction(s string) int {
	chars := strings.Split(s, "")

	result1 := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i, _ := range chars {
		result2 := stringReductionProcess(s, i)

		if len(result1) > len(result2) {
			result1 = result2
		}
	}

	return len(result1)
}

func main() {
	var strings []string = []string{"cab", "bcab", "ccccc", "abcbcba", "ababbac", "abababaccc", "cacacacbcbcaccccc"}

	// strings := []string{
	// 	"cacacacbcbcaccccc",
	// }

	for _, s := range strings {
		fmt.Println("Length ========", stringReduction(s))
	}
	// 2
	// 1
	// 5

	// 1
	// 2
	// 1

	// 3
}
