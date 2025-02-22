package main

import (
    "fmt"
    "time"
)

func main() {
    getTime()
}
func getTime() {
    t := time.Now()
    fmt.Println(t)
}
