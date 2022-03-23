package main

// https://www.hackerrank.com/challenges/the-hurdle-race/problem?isFullScreen=true
func hurdleRace(k int32, height []int32) int32 {
    sort.Slice(height, func(i, j int) bool { return height[i] < height[j] })

	dose := height[len(height)-1] - k
	if dose < 0{
		return 0
	}

    return dose
}
