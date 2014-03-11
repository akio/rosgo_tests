package main

import (
    "fmt"
    "ros"
    "rosgo_tests"
)


func callback(msg *rosgo_tests.Hello) {
    fmt.Printf("Received: %s\n", msg.Data)
}


func main() {
    node := ros.NewNode("/listener")
    defer node.Shutdown()
    node.Logger().SetSeverity(ros.LogLevelDebug)
    node.NewSubscriber("/chatter", rosgo_tests.MsgHello, callback)
    node.Spin()
}
