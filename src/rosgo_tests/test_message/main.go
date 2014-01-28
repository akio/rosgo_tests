package main

import (
//    "io"
//    "bytes"
    "fmt"
    "testing"
    rosgo_tests "rosgo_tests/gen"
)


func TestSerialize(t *testing.T) {
    msg := rosgo_tests.AllFieldTypes{}
    fmt.Println(msg)
}

func TestDeserialize(t *testing.T) {
    msg := rosgo_tests.AllFieldTypes{}
    fmt.Println(msg)
}


func main() {
    fmt.Println("Hello")
    msg := rosgo_tests.AllFieldTypes{}
    fmt.Println(msg.H)
    fmt.Println(msg.B)
    fmt.Println(msg.I8)
    fmt.Println(msg.I16)
    fmt.Println(msg.I32)
    fmt.Println(msg.I64)
    fmt.Println(msg.U8)
    fmt.Println(msg.U16)
    fmt.Println(msg.U32)
    fmt.Println(msg.U64)
    fmt.Println(msg.F32)
    fmt.Println(msg.F64)
    fmt.Println(msg.T)
    fmt.Println(msg.D)
    fmt.Println(msg.S)
    fmt.Println(msg.C)
    fmt.Println(msg.DynAry)
    fmt.Println(msg.FixAry)
}
