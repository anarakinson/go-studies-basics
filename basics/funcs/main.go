package main

import (
    "fmt"
)

func contains(str_arr []string, str string) bool {
    res := false
    for _, v := range str_arr {
        if v == str {
            res = true
        }
    }
    return res
}

func getMax(arr ... int) int {
    var res int
    for _, v := range arr {
        if v > res {
            res = v
        }
    }
    return res
}

func main() {
    var str = "abra"
    var str1 = "uuu"
    str_array := []string{
        "abra", "cadabra", "alacazam",
    }

    fmt.Println(contains(str_array, str))
    fmt.Println(contains(str_array, str1))
    fmt.Println(getMax(123, 2, 3, 41, 5))
}
