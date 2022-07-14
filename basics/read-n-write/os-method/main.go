package main

import (
    "fmt"
    "os"
    "io"
)

func getFile(path string) []byte {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Cannot open:", path)
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    var data []byte
    buf := make([]byte, 1024)

    for {
        n, err := file.Read(buf)
        if (err != io.EOF) && (err != nil) {
            panic(err)
        }
        if n == 0 {
            break
        }
        data = append(data, buf[:n]...)
    }
    // fmt.Println(string(data))
    return data
}

func writeFile(data []byte, filepath string) {
    file, err := os.Create(filepath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    _, err = file.Write(data) // for bytes
    // _, err := file.WriteString(data) // for strings
    if err != nil {
        panic(err)
    }
    file.Sync()
}

func main() {
    path := "../test.txt"
    data := getFile(path)
    fmt.Println(string(data))
    writeFile(data, "out.txt")
}
