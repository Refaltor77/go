package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup 

func run(name string) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        time.Sleep(1 * time.Second)
        fmt.Println(name, " : ", i)
    }
}

func main() {
    debut := time.Now()

    wg.Add(1)
    go run("Refaltor")
    wg.Add(1)
    go run("Natof")
    wg.Add(1)
    go run("Flolm")

    wg.Wait()
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
}
