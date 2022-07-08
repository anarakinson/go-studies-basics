package main

import (
    "fmt"
)

func main() {
    var a *int
    var b int = 12
    fmt.Println("a - ", a, "\nb - ", b, "\n")

    a = &b
    *a = 13
    fmt.Println("a - ", *a, "\nb - ", b, "\n")

}
