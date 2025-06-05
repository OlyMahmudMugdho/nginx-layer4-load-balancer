package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    hostname, _ := os.Hostname()
    listener, err := net.Listen("tcp", ":5000")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Printf("TCP Server running on %s\n", hostname)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go func(c net.Conn) {
            defer c.Close()
            msg := fmt.Sprintf("sending hello from %s\n", hostname)
            c.Write([]byte(msg))
        }(conn)
    }
}

