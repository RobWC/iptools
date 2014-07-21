package main

import "net"
import "fmt"


func main() {
    testIP := net.ParseIP("255.0.0.0")
    fmt.Println([]byte(testIP))
    //fmt.Println([]byte())
    fmt.Println([]byte{0,0xff,0,0,0})
}
