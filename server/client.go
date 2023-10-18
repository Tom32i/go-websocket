package server

import (
    "log"
    // "encoding/json"
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
		/*

		jsonErr := json.Unmarshal(data, &message)
		if jsonErr != nil {
			log.Printf("error decoding json: %v", jsonErr)
			break
		}*/

		server.in <- message
	}
}

// func parseMessage(message Message) Message {
// 	switch m.name {
// 		case "name":
// 			return ClientNameMessage{
// 				client: c,
// 				name: m.name,
// 				data: interface{}{
// 					name: m.data,
// 					id: client.id,
// 				},
// 			}
// 		default:
// 			return message
// 	}
// }



