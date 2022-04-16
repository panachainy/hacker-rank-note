package main

// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem?isFullScreen=true
func invertInt(i int32) int32{
    inverted := ""
    iStr := strconv.Itoa(int(i))
    for in := len(iStr)-1; in >= 0; in-- {
        inverted += string(iStr[in])
    }
    result, _ := strconv.Atoi(inverted)
    return int32(result)
}

func beautifulDays(i int32, j int32, k int32) int32 {
    butifulCount := int32(0)
    for s:=i; s<=j; s++ {
        invertS := invertInt(s)

        result := math.Abs(float64(s-invertS))/float64(k)
        resultStr := fmt.Sprint(result)
        if _, err := strconv.Atoi(resultStr); err == nil {
            butifulCount++
        }
    }
    return butifulCount
}
