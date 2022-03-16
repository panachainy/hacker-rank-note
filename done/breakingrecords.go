package main

// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem?isFullScreen=false
func breakingRecords(scores []int32) []int32 {
	result := []int32{0, 0}

	var min int32 = scores[0]
	var hight int32 = scores[0]

	for i := 1; i < len(scores); i++ {
		if min > scores[i] {
			min = scores[i]
			result[1]++
		}

		if hight < scores[i] {
			hight = scores[i]
			result[0]++
		}
	}

	return result
}
