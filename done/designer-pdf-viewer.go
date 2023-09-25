package main

import (
	"fmt"
	"strings"
)

func letterToPosition(letter string) int32 {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	letter = strings.ToLower(letter)
	position := strings.Index(alphabet, letter)
	return int32(position)
}

func designerPdfViewer(h []int32, word string) int32 {
	highest := int32(0)
	characters := strings.Split(word, "")
	for _, character := range characters {
		charIndex := letterToPosition(character)
		fmt.Println(h[charIndex])

		if highest <= h[charIndex] {
			highest = h[charIndex]
		}
	}

	result := len(word) * 1 * int(highest)
	return int32(result)
}

func main() {
	result := designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "zaba")
	fmt.Println("res: ", result)

}
