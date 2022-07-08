package main

import (
    "fmt"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
    var speed float64 = 120.4 // m/sec
    european := EuropeanVelocity((speed / 1000) * 3600)
    american := AmericanVelocity((speed / 1609) * 3600)

    fmt.Println(fmt.Sprintf("%.2f m/s", speed))
    fmt.Println(fmt.Sprintf("%.2f km/h", european))
    fmt.Println(fmt.Sprintf("%.2f mph", american))
}
