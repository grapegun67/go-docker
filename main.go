package main

import (
    "fmt"
    "os"
)

func main() {
    file1, _ := os.Create("/tmp/test.txt")
    defer file1.Close()
    fmt.Fprint(file1, 1, 1.1, "Success")
}
