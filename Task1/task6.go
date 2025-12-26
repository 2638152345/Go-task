package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    for i := 1; i < len(nums); {
        if nums[i] == nums[i-1] {
            for j := i; j < len(nums)-1; j++ {
                nums[j] = nums[j+1]
            }
            nums = nums[:len(nums)-1]
        } else {
            i++
        }
    }
    return len(nums)
}

func main() {
    var n int
    fmt.Scan(&n)
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
    }

    k := removeDuplicates(nums)
    fmt.Println(k)
    for i := 0; i < k; i++ {
        fmt.Print(nums[i], " ")
    }
}
