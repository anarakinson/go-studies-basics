package main

import (
    "fmt"
)

func main() {
    var week = []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
    workdays := make([]string, 5)
    copy(workdays, week[1:len(week)])
    week = append(week[:1], week[len(week)-1:]...)

    fmt.Println("holidays - ", week)
    fmt.Println("workdays - ", workdays)
}
