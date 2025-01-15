package main

import (
        "fmt"
        "sync"
)

func main() {

        fmt.Println("Corrected Concurrency in Go")
        var wg sync.WaitGroup
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("Goroutine %d: %d\n", i, i)
                }(i)
        }
        wg.Wait()
}