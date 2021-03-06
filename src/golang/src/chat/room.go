package main

import (
	"log"
	"net/http"

	"github.com/gorakhargosh/trace"
	"github.com/gorilla/websocket"
)

const (
	connectionBufferSize = 1024
	messageBufferSize    = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  connectionBufferSize,
	WriteBufferSize: connectionBufferSize,
}

type room struct {
	// forward is a channel that holds incoming messages that should be
	// forwarded to other clients.
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	tracer  trace.Tracer
}

// Creates an new instance of a room.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Discard(),
	}
}

// Kick starts a room and blocks for a conversation.
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("New client joined.")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left.")
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", string(msg))
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
					r.tracer.Trace(" -- sent to client.")
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- failed to send, cleaned up client.")
				}
			}
		}
	}
}

// The room is an HTTP handler.
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		conn: conn,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}
	r.join <- client
	defer func() { r.leave <- client }() // If the connection closes, clean up.
	go client.write()                    // Keep writing to the client in a goroutine.
	client.read()                        // Block reading.
}
