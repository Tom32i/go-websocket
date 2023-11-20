package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "curvygo/server"
)

func main() {
    port := flag.Int("port", 8034, "Port to run on")
    public := flag.String("public", "./public", "Public path to serve")

    flag.Parse()

    server := server.CreateServer()

    if *public != "" {
        http.Handle("/", http.FileServer(http.Dir(*public)))
    }

    http.HandleFunc("/ws", server.Handler)

    log.Printf("Launching server")
    go server.Run()

    log.Printf("Listening on port %d", *port)
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
