package main

import (
    "fmt"
    "io/ioutil"
)

func Read() ([]byte) {
    data, err := ioutil.ReadFile("../test.txt")
    if err != nil {
        fmt.Println("Cannot read file:\n", err)
    }
    return data
}

func Write(path string, data []byte) {
    err := ioutil.WriteFile(path, data, 0666)
    if err != nil {
        fmt.Println("Cannot write file:\n", err)
    }
}

func main() {
    data := Read()
    print(string(data))
    Write("out.txt", data)
}
