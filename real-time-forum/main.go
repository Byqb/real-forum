package main

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    _ "github.com/mattn/go-sqlite3"
)

var upgrader = websocket.Upgrader{}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/posts", postsHandler)
    http.HandleFunc("/post", postHandler)
    http.HandleFunc("/comment", commentHandler)
    http.HandleFunc("/messages", messagesHandler)
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// WebSocket handler
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error while connecting:", err)
        return
    }
    // Handle WebSocket connection
}

// Other handlers like registerHandler, loginHandler, etc.
