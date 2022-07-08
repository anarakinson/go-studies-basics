package main

import (
    "fmt"
    "math"
)

func main() {
    m := map[string]float32{"first": 1, "second": 2, "third": 3}
    fmt.Println(m)
    
    x, ok := m["fourth"]
    if !ok {
        fmt.Println(x, ok)
    }

    delete(m, "second")

    fmt.Println(m)

    m["fifth"] = 12
    fmt.Println(m)

    var a string = "hello"
    fmt.Println(a)
    var b = []rune(a)
    b[2] = '!'
    var c = string(b)
    fmt.Println(c)

    fmt.Println(fmt.Sprintf("%.2f", math.Round(9.7)))
    fmt.Println(math.Pow(2, 2))
}
