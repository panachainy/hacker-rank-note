package main

// https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true
func sockMerchant(n int32, ar []int32) int32 {
    sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })

    currentSock := int32(0)
    countPairedSock := int32(0)

    for _, v := range ar{
        if currentSock == 0 {
            currentSock = v
            continue
        }

        if v == currentSock{
            countPairedSock++
            currentSock = 0
        }else {
            currentSock = v
        }
    }

    return countPairedSock
}
