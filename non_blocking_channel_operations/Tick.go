package main

import (
    "fmt"
    "time"
)

func main() { 
    limiter := time.Tick(500 * time.Millisecond)

    for i := 1; i <= 10; i++ {
        <-limiter
        fmt.Println("tick", time.Now())
    } 
}


