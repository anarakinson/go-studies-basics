package main

import (
    "fmt"
)

func increment() func() int {
    i := 0
    return func() int {
        i = i + 1
        return i
    }
}

func main() {
    var i = 0
    for {
        i = i + 1
        fmt.Println(i)
        if i >= 5 {
            break
        }
    }

    fmt.Println()

    inc := increment()

    for i := 0; i < 5; i++ {
        fmt.Println(inc())
    }

}
