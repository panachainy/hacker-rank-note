package main

// https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true
func birthday(s []int32, d int32, m int32) int32 {
	dividedWay := int32(0)

	if len(s) == 1 && s[0] == d {
		return 1
	}

	for i := 0; i < len(s); i++ {
		// check if can't sum loop with j+m should break;
		if int32(i)+m > int32(len(s)) {
			break
		}

		sum := int32(0)
		for x := 0; x < int(m); x++ {
			sum += s[i+x]
		}

		if sum == d {
			dividedWay++
		}
	}

	return dividedWay
}
