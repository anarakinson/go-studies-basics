package main

import (
    "fmt"
    "go-studies/basics/packages/words"
    _ "go-studies/basics/packages/init"
)

func main() {
    fmt.Println(words.Hello)
    for i := 0; i < 5; i++ {
        fmt.Println(words.Random())
    }
}
