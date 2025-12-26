package main

import "fmt"

func addTen(x *int) {
    *x = *x + 10
}

func main() {
    var num int
    fmt.Scan(&num)

    addTen(&num)

    fmt.Println(num)
}
