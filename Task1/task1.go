package main

import "fmt"

func singleNumber(nums []int) int {
    for i := 0; i < len(nums); i++ {
        count := 0
        for j := 0; j < len(nums); j++ {
            if nums[i] == nums[j] {
                count++
            }
        }
        if count == 1 {
            return nums[i]
        }
    }
    return 0
}

func main() {
    var n int
    fmt.Scan(&n)
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
    }
    fmt.Println(singleNumber(nums))
}

