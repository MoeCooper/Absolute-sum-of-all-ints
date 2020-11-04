package main

import (
    "fmt"
    "math"
)

func sumNums(arr[] int) int64 {
    var sum int64 = 0
    for i := 0; i < len(arr); i++ {
        sum += int64(math.Abs(float64(arr[i])))
    }
    return sum 
}

func main() {
    myArr := []int{1,2,-3,4,5}
    fmt.Println(sumNums(myArr))
}