package main

import "fmt"

func appendAndDelete(s string, t string, k int32) string {
	primary := ""
	secondary := ""

	if len(s) > len(t) {
		primary = s
		secondary = t
	} else {
		primary = t
		secondary = s
	}

	diffIndex := 0

	primaryLength := len(primary)
	secondaryLength := len(secondary)

	for i := 0; i < primaryLength; i++ {
		if secondaryLength < i+1 {
			diffIndex = i
			break
		}

		if primary[i] != secondary[i] {
			diffIndex = i
			break
		}
	}

	if diffIndex == 0 {
		return "Yes"
	}

	amountAllDiff := int32(primaryLength - diffIndex + secondaryLength - diffIndex)

	if k == amountAllDiff || k >= int32(primaryLength+secondaryLength) {
		return "Yes"
	} else if amountAllDiff%2 == k%2 && amountAllDiff <= k {
		return "Yes"
	}

	return "No"
}

func main() {
	fmt.Println("expect - yes", "result:",
		appendAndDelete("hackerhappy", "hackerrank", 9))

	fmt.Println("==")

	fmt.Println("expect - yes", "result:",
		appendAndDelete("aaa", "a", 5))

	fmt.Println("==")

	fmt.Println("expect - no", "result:",
		appendAndDelete("ashley", "ash", 2))
}
