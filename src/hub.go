package src

import (
	"encoding/json"
	"fmt"
)

const (
	formatJoin  = "New client joins room (ID: %s)"
	formatLeave = "A client leaves room (ID: %s)"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			clientId := client.ID
			str := fmt.Sprintf(formatJoin, clientId)
			for client := range h.clients {
				msg := []byte(str)
				client.send <- msg
			}

			h.clients[client] = true

		case client := <-h.unregister:
			clientId := client.ID
			str := fmt.Sprintf(formatLeave, clientId)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			for client := range h.clients {
				msg := []byte(str)
				client.send <- msg
			}
		case userMessage := <-h.broadcast:
			var data map[string][]byte
			_ = json.Unmarshal(userMessage, &data)

			for client := range h.clients {
				//prevent self receive the message
				if client.ID == string(data["id"]) {
					continue
				}
				select {
				case client.send <- data["message"]:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
