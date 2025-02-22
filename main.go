package main

import (
    "fmt"
    "time"
)

func main() {
    getTime()
    new1()
}
func getTime() {
    t := time.Now()
    fmt.Println(t)
}
