package main

import "fmt"

func multiplyByTwo(nums *[]int) {
    for i := 0; i < len(*nums); i++ {
        (*nums)[i] = (*nums)[i] * 2
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
    }

    multiplyByTwo(&nums)

    for i := 0; i < len(nums); i++ {
        fmt.Print(nums[i], " ")
    }
}
