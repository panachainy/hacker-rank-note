package main

import "fmt"

func minLength(s string, t string) string {
	if len(s) < len(t) {
		return s
	}

	return t
}

// https://www.hackerrank.com/challenges/append-and-delete/problem
// from leaderboard
func appendAndDelete(s string, t string, k int32) string {
	sameLength := int32(0)

	sLength := int32(len(s))
	tLength := int32(len(t))

	for i := 0; i < len(minLength(s, t)); i++ {
		if s[i] == t[i] {
			sameLength++
		} else {
			break
		}
	}

	if (sLength + tLength - (2 * sameLength)) > k {
		return "No"
	}

	if (sLength+tLength-(2*sameLength))%2 == k%2 {
		return "Yes"
	}

	if (sLength + tLength - k) < 0 {
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
