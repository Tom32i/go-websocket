package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "curvygo/server"
)

func main() {
    port := flag.Int("port", 8032, "Port to run on")

    flag.Parse()

    server := server.CreateServer()

    http.Handle("/", http.FileServer(http.Dir("./public")))
    http.HandleFunc("/ws", server.Handler)

    log.Printf("Launching server")
    go server.Run()

    log.Printf("Listening on port %d", *port)
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
