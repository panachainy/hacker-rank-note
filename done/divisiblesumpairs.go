package main

// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var divisibleCount int32 = 0

	for i, vi := range ar {
		var j int32 = 0
		for j = int32(i + 1); j < n; j++ {
			sum := vi + ar[j]
			if sum%k == 0 {
				divisibleCount++
			}
		}
	}

	return divisibleCount
}
