package main

import (
    "fmt"
    "io"
    "net"
    "net/http"
)

func tcpFetchFromNginx() string {
    conn, err := net.Dial("tcp", "localhost:5000")
    if err != nil {
        return fmt.Sprintf("Error connecting: %v", err)
    }
    defer conn.Close()

    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil && err != io.EOF {
        return fmt.Sprintf("Error reading: %v", err)
    }

    return string(buf[:n])
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        msg := tcpFetchFromNginx()
        w.Write([]byte(msg))
    })

    fmt.Println("HTTP server running on :8080")
    http.ListenAndServe(":8080", mux)
}

