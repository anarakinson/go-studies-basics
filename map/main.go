package main

import (
    "fmt"
    "sort"
)

func main() {
    var libr = map[string]map[string]string {
        "first author": {
            "name": "Stephen",
            "lastname": "King",
            "book": "It",
        },
        "second author": {
            "name": "Lev",
            "lastname": "Tolstoy",
            "book": "War and Peace",
        },
        "third author": {
            "name": "Hovard",
            "lastname": "Lovecraft",
            "book": "Call of Chtulhu",
        },
        "fourth author": {
            "name": "William",
            "lastname": "Shakespear",
            "book": "Hamlet",
        },
    }

    first, ok := libr["first author"]["book"]
    if ok {
        fmt.Println(first)
    }
    fmt.Println(len(libr))

    libr_keys := make([]string, 0, len(libr))
    for key := range(libr) {
        libr_keys = append(libr_keys, key)
    }
    sort.Strings(libr_keys)

    for _, v := range libr_keys {
        fmt.Println(v, "-", libr[v]["book"])
    }
}
