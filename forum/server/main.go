package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Структура для сообщения
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var (
	messages = []Message{}
	idCount  = 0
	mutex    = sync.Mutex{}
)

func main() {
	http.HandleFunc("/api/messages", handleMessages)
	http.Handle("/", http.FileServer(http.Dir("../client")))

	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Обработчик для сообщений
func handleMessages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	case "POST":
		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)

		mutex.Lock()
		idCount++
		msg.ID = idCount
		messages = append(messages, msg)
		mutex.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
