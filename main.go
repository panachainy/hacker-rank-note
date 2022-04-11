package main

import "fmt"

func appendAndDelete(s string, t string, k int32) string {
	fmt.Println("s", s)
	fmt.Println("t", t)

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
		// fmt.Println("check not equal",primary[i],secondary[i])
		// fmt.Println("diffIndex",diffIndex)

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

	// kIsEven := false

	// if k % 2 == 0 {
	//     kIsEven = true
	// }

	// if int32(amountAllDiff) > k {
	//     return "No"
	// }

	// allDiffIsEven := false
	// if int32(amountAllDiff) % 2 == 0 {
	//     allDiffIsEven = true
	// }

	// if kIsEven == allDiffIsEven {
	//     return "Yes"
	// }

	// return "No"

	if k == amountAllDiff || k >= int32(primaryLength+secondaryLength) {
		return "Yes"
	} else if amountAllDiff%2 == k%2 && amountAllDiff <= k {
		return "Yes"
	}

	return "No"
}

func main() {
	resp := appendAndDelete("hackerhappy", "hackerrank", 9)
	fmt.Println("result:", resp)
}
