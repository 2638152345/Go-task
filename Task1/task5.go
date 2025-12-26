package main

import "fmt"

func plusOne(digits []int) []int {
    n := len(digits)
    for i := n - 1; i >= 0; i-- {
        digits[i]++
        if digits[i] < 10 {
            return digits
        }
        digits[i] = 0
    }
    newDigits := make([]int, n+1)
    newDigits[0] = 1
    return newDigits
}

func main() {
    var n int
    fmt.Scan(&n)
    digits := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&digits[i])
    }

    result := plusOne(digits)
    for _, v := range result {
        fmt.Print(v, " ")
    }
}
