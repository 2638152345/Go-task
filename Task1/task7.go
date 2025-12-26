package main

import "fmt"

func merge(intervals [][]int) [][]int {
    n := len(intervals)
    if n == 0 {
        return [][]int{}
    }

    for i := 0; i < n-1; i++ {
        for j := 0; j < n-1-i; j++ {
            if intervals[j][0] > intervals[j+1][0] {
                intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
            }
        }
    }

    merged := [][]int{intervals[0]}
    for i := 1; i < n; i++ {
        last := merged[len(merged)-1]
        if intervals[i][0] <= last[1] {
            if intervals[i][1] > last[1] {
                last[1] = intervals[i][1]
            }
        } else {
            merged = append(merged, intervals[i])
        }
    }

    return merged
}

func main() {
    var n int
    fmt.Scan(&n)
    intervals := make([][]int, n)
    for i := 0; i < n; i++ {
        intervals[i] = make([]int, 2)
        fmt.Scan(&intervals[i][0], &intervals[i][1])
    }

    result := merge(intervals)
    for _, interval := range result {
        fmt.Println(interval[0], interval[1])
    }
}
