package services

// Hub maintains the set of active ChatConnections and broadcasts messages to the
// ChatConnections.
type Hub struct {
	// Registered ChatConnections.
	ChatConnections map[*ChatConnection]bool

	// Inbound messages from the ChatConnections.
	Broadcast chan []byte

	// Register requests from the ChatConnections.
	Register chan *ChatConnection

	// Unregister requests from ChatConnections.
	Unregister chan *ChatConnection
}

var Hubs = Hub{
	Broadcast:   make(chan []byte),
	Register:    make(chan *ChatConnection),
	Unregister:  make(chan *ChatConnection),
	ChatConnections: make(map[*ChatConnection]bool),
}

func (hub *Hub) Run() {
	for {
		select {
		case c := <-hub.Register:
			hub.ChatConnections[c] = true
		case c := <-hub.Unregister:
			if _, ok := hub.ChatConnections[c]; ok {
				delete(hub.ChatConnections, c)
				close(c.Send)
			}
		case m := <-hub.Broadcast:
			for c := range hub.ChatConnections {
				select {
				case c.Send <- m:
				default:
					close(c.Send)
					delete(hub.ChatConnections, c)
				}
			}
		}
	}
}
