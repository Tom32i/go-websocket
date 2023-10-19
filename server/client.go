package server

import (
    "log"
    "github.com/gorilla/websocket"
)

type Client struct {
	id uint8
	name string
	socket *websocket.Conn
    encoder BinaryEncoder
}

func (c *Client) setName(name string) {
	c.name = name
}

func (c *Client) write(data []byte) {
	c.socket.WriteMessage(websocket.BinaryMessage, data)
}

func (c *Client) run(server *Server) {
	defer func() {
		server.removeClient(c)
		c.socket.Close()
	}()

	//c.socket.SetReadLimit(maxMessageSize)
	//c.socket.SetReadDeadline(time.Now().Add(pongWait))
	//c.socket.SetPongHandler(func(string) error {
	//	c.socket.SetReadDeadline(time.Now().Add(pongWait)); return nil
	//})

	for {
		_, data, err := c.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message := c.encoder.decode(data)
		message.client = c

		server.in <- message
	}
}
