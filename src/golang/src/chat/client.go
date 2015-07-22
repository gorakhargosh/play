package main

import "github.com/gorilla/websocket"

// import "golang.org/x/net/websocket"

// client represents a single chatting connection.
type client struct {
	conn *websocket.Conn
	send chan []byte
	room *room
}

// Blocking read from the conn.
func (c *client) read() {
	for {
		if _, msg, err := c.conn.ReadMessage(); err != nil {
			break
		} else {
			c.room.forward <- msg
		}
	}
	c.conn.Close()
}

// Blocking write to the conn.
func (c *client) write() {
	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.conn.Close()
}
