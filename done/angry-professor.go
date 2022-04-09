package main

// https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true
func angryProfessor(k int32, a []int32) string {
	onTimeCount := int32(0)

	for _, v := range a {
		if v <= 0 {
			onTimeCount++
		}
	}

	if k > onTimeCount {
		return "YES"
	}

	return "NO"
}
