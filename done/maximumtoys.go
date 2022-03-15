package main

import "sort"

// https://www.hackerrank.com/challenges/mark-and-toys/problem?isFullScreen=true
func maximumToys(prices []int32, k int32) int32 {
	// Sort less to more
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })

	var toyCount int32 = 0
	for _, v := range prices {
		if k >= v {
			k -= v
			toyCount++
		} else {
			break
		}
	}

	return toyCount
}
