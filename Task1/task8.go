package main

import "fmt"

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}

func main() {
    var n int
    fmt.Scan(&n)

    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
    }

    var target int
    fmt.Scan(&target)

    result := twoSum(nums, target)
    for i := 0; i < len(result); i++ {
        fmt.Print(result[i], " ")
    }
}
