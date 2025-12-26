package main

import (
    "fmt"
    "sync"
    "time"
)

func runTask(id int, task func(), wg *sync.WaitGroup) {
    defer wg.Done()
    start := time.Now()
    task()
    duration := time.Since(start)
    fmt.Println("任务", id, "执行时间:", duration)
}

func main() {
    tasks := []func(){
        func() {
            time.Sleep(1 * time.Second)
        },
        func() {
            time.Sleep(500 * time.Millisecond)
        },
        func() {
            time.Sleep(2 * time.Second)
        },
    }

    var wg sync.WaitGroup
    wg.Add(len(tasks))

    for i := 0; i < len(tasks); i++ {
        go runTask(i+1, tasks[i], &wg)
    }

    wg.Wait()
}
