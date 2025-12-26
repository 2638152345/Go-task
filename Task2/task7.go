package main

import (
    "fmt"
    "sync"
)

func producer(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 10; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for v := range ch {
        fmt.Println(v)
    }
}

func main() {
    ch := make(chan int)

    var wg sync.WaitGroup
    wg.Add(2)

    go producer(ch, &wg)
    go consumer(ch, &wg)

    wg.Wait()
}
