package main

import (
  "log"
  "io"  
  "net/http"
  "golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
  io.Copy(ws, ws)
}

func main() {
  // REST hello world
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Println("Got request")
    w.Write([]byte("Hello world"))
  })
  
  // Websocket echo
  http.Handle("/echo", websocket.Handler(EchoServer))
  
  // Start server 
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Println("ListenAndServe:", err)
  }
}

