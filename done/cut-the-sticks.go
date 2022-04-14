package main

// https://www.hackerrank.com/challenges/cut-the-sticks/problem?isFullScreen=true
func getLengthNotZero(arr []int32) int32{
    count := int32(0)
    for _, v := range arr{
        if v >= 0 {
            count++
        }
    }
    return count
}

func getMinNotZero(arr []int32) int32{
    min := int32(math.MaxInt32)
    for _, v := range arr{
        if v > 0 && v < min{
            min = v
        }
    }
    
    if min == int32(math.MaxInt32){
        return 0
    }
    
    return min
}


func cutTheSticks(arr []int32) []int32 {
    result := []int32{}

    for i := 0; i<len(arr);i++ {
        min := getMinNotZero(arr)
        
        if min == 0 {
            break;
        }
        
        for j := range arr {
            arr[j] -= min
        }
        
        result = append(result, getLengthNotZero(arr))
    }
  

    return result
}
