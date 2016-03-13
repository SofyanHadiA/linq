package services

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10

	// Maximum message size allowed from peer.
	MaxMessageSize = 512
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ChatConnection is an middleman between the websocket ChatConnection and the hub.
type ChatConnection struct {
	// The websocket ChatConnection.
	WS *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket ChatConnection to the hub.
func (c *ChatConnection) ReadPump() {
	defer func() {
		 Hubs.Unregister <- c
		c.WS.Close()
	}()
	c.WS.SetReadLimit(MaxMessageSize)
	c.WS.SetReadDeadline(time.Now().Add(PongWait))
	c.WS.SetPongHandler(func(string) error { c.WS.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		 Hubs.Broadcast <- message
	}
}

// write writes a message with the given message type and payload.
func (c *ChatConnection) Write(mt int, payload []byte) error {
	c.WS.SetWriteDeadline(time.Now().Add(WriteWait))
	return c.WS.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket ChatConnection.
func (c *ChatConnection) WritePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		c.WS.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

