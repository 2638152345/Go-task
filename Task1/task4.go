package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        for len(prefix) > 0 && (len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix) {
            prefix = prefix[:len(prefix)-1]
        }
        if len(prefix) == 0 {
            return ""
        }
    }

    return prefix
}

func main() {
    var n int
    fmt.Scan(&n)
    strs := make([]string, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&strs[i])
    }
    fmt.Println(longestCommonPrefix(strs))
}
