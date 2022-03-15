package drawingbook

// https://www.hackerrank.com/challenges/drawing-book/problem?isFullScreen=true
func pageCount(n int32, p int32) int32 {
    var result int32 = 0
    var forward int32 = p/2
    var backward int32 = 0
    
    // backward logic is not cool TT
    // before I use `(n-p)/2` but stuck in case n=6 p=5 (จำโจทไม่ได้ แต่ข้อที่ 26)
    
    startLoop := n-1
    
    if n % 2 == 0{
        startLoop = n
    }
    
    for i := startLoop; i >= p; i-- {
       if i % 2 == 1{
           backward++
       }
    }
    
    if forward < backward {
        result = forward
    }else{
        result = backward
    }
    
    return result
}
