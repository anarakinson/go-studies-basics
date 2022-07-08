package main

import (
    "fmt"
)

func main() {
    var AmericanVelocity, EuropeanVelocity float64
    var speed float64 = 120.4 // m/sec
    EuropeanVelocity = (speed / 1000) * 3600
    AmericanVelocity = (speed / 1609) * 3600
    fmt.Println(fmt.Sprintf("%.2f m/s", speed))
    fmt.Println(fmt.Sprintf("%.2f km/h", EuropeanVelocity))
    fmt.Println(fmt.Sprintf("%.2f mph", AmericanVelocity))
}
