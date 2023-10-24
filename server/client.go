package server

import (
    "log"
    "github.com/gorilla/websocket"
    "curvygo/server/codec"
)

type Client struct {
	id uint8
	name string
	socket *websocket.Conn
    encoder codec.BinaryEncoder
	x uint16
	y uint16
}

type ClientMessage struct {
    client *Client
    message codec.Message
}

func (c *Client) setName(name string) {
	c.name = name
	log.Printf("Client #%v name is '%v'.", c.id, c.name)
}

func (c *Client) setPosition(x uint16, y uint16) {
	c.x = x
	c.y = y
}

func (c *Client) write(data []byte) {
	c.socket.WriteMessage(websocket.BinaryMessage, data)
}

func (c *Client) run(server *Server) {
	defer func() {
		server.removeClient(c)
		c.socket.Close()
	}()

	for {
		_, data, err := c.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		server.in <- ClientMessage{c, c.encoder.Decode(data)}
	}
}
