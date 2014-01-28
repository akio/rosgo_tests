package main

import (
    "fmt"
    "time"
    "ros"
    rosgo_tests "rosgo_tests/gen"
)


func main() {
    node := ros.NewNode("/talker")
    defer node.Shutdown()
    node.Logger().SetSeverity(ros.LogLevelDebug)
    pub := node.NewPublisher("/chatter", rosgo_tests.MsgHello)

    for node.OK() {
        node.SpinOnce()
        var msg rosgo_tests.Hello
        msg.Data = fmt.Sprintf("hello %s", time.Now().String())
        fmt.Println(msg.Data)
        pub.Publish(&msg)
        time.Sleep(time.Second)
    }
}
