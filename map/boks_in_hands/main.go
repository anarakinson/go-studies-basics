package main

import (
    "fmt"
)

func main() {
    onHands := map[string][]string {}
    onHands["Billy"] = []string{"1 book", "4 book"}
    onHands["Bobby"] = []string{"2 book", "3 book", "6 book", "10 book"}
    onHands["Jimmy"] = []string{"5 book", "7 book", "9 book"}

    for k, v := range onHands {
        fmt.Println(k, v)
    }
}
