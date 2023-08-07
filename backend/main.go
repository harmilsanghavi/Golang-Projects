package main

import (
	"chatreactgo/backend/pkg/websocket"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Chat Application")
	setRoutes()
}

func setRoutes() {
	http.HandleFunc("/", DefaultLoad)
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
	http.ListenAndServe(":8080", nil)
}

func DefaultLoad(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple Server")
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
