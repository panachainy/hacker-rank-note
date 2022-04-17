package main

import (
	"fmt"
)

// https://www.hackerrank.com/challenges/repeated-string/problem?isFullScreen=true
func repeatedString(s string, n int64) int64 {
	stringLength := int64(len(s))
	aInOneRound := int64(0)

	// Count a in one round
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" {
			aInOneRound++
		}
	}

	round := n / stringLength

	// not include scrap
	result := aInOneRound * round

	currentCount := round * stringLength

	// Calculate scrap
	currentIndex := 0
	for i := currentCount; i < n; i++ {
		if string(s[currentIndex]) == "a" {
			result++
		}
		currentIndex++
	}

	return result
}

func main() {
	fmt.Println("expect - 16481469408", "result:",
		repeatedString("epsxyyflvrrrxzvnoenvpegvuonodjoxfwdmcvwctmekpsnamchznsoxaklzjgrqruyzavshfbmuhdwwmpbkwcuomqhiyvuztwvq", 549382313570))

	fmt.Println("===========")

	fmt.Println("expect - 7", "result:",
		repeatedString("aba", 10))

	fmt.Println("===========")

	fmt.Println("expect - 1000000000000", "result:",
		repeatedString("a", 1000000000000))
}
