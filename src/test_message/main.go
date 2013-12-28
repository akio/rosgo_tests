package main

import (
    "fmt"
    "rosgo_tests"
)

func main() {
    fmt.Println("Hello")   
    msg := rosgo_tests.MsgAllFieldTypes.NewMessage()
}
