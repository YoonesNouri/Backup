package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var clientsMutex sync.Mutex

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	go broadcastLoop()

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	// Read messages (you can implement your own logic here)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
	}

	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()

	fmt.Println("Client disconnected")
}

func broadcastLoop() {
	for {
		// Simulate a backend change
		message := "Backend change occurred: " + time.Now().Format(time.RFC3339)

		clientsMutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				fmt.Println("Error sending message:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMutex.Unlock()

		// Wait for a while before sending the next update
		time.Sleep(5 * time.Second)
	}
}
