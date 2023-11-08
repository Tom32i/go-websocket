package server

import (
	"curvygo/server/codec"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Server struct {
	upgrader websocket.Upgrader
	clients  map[uint8]*Client
	in       chan ClientMessage
	encoder  codec.BinaryEncoder
}

func (server *Server) Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Upgrading connexion")

	socket, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client, error := server.createClient(socket)

	if error != nil {
		log.Println(error)
		return
	}

	server.init(client)

	go client.run(server)
}

func (server *Server) Run() {
	log.Printf("Server running")
	for {
		m := <-server.in
		//log.Printf("message in: %v (%v)", m.message.Name, m.message.Data)
		switch m.message.Name {
		case "me:name":
			m.client.setName(m.message.Data.(string))
			server.writeAll(
				"client:name",
				codec.ClientNameMessage{m.client.id, m.client.name},
			)
		case "me:position":
			position := m.message.Data.(codec.Position)
			m.client.setPosition(position.X, position.Y)
			server.writeAll(
				"client:position",
				codec.ClientPosition{m.client.id, codec.Position{m.client.x, m.client.y}},
			)
		}
	}
}

func (server *Server) createClient(socket *websocket.Conn) (*Client, error) {
	id, error := server.nextId()

	if error != nil {
		return nil, error
	}

	c := Client{
		id:      id,
		socket:  socket,
		encoder: server.encoder,
	}

	server.clients[c.id] = &c

	log.Printf("Client #%d joined.", c.id)

	return &c, nil
}

func (server *Server) removeClient(c *Client) {
	delete(server.clients, c.id)
	server.writeAll("client:remove", c.id)
	log.Printf("Client #%d left.", c.id)
}

func (server Server) writeAll(name string, data any) {
	buffer := server.encoder.Encode(name, data)
	for _, c := range server.clients {
		c.write(buffer)
	}
}

func (server *Server) init(client *Client) {
	// Send the client to everybody
	server.writeAll("client:add", codec.ClientAddMessage{client.id, client.name})

	// Send the client its id
	client.write(server.encoder.Encode("me:id", client.id))

	// Send the clients the current client list and positions
	for _, c := range server.clients {
		if c.id != client.id {
			client.write(server.encoder.Encode("client:add", codec.ClientAddMessage{c.id, c.name}))
			client.write(server.encoder.Encode("client:position", codec.ClientPosition{c.id, codec.Position{c.x, c.y}}))
		}
	}
}

func (server *Server) nextId() (uint8, error) {
	for i := 0; i < 256; i++ {
		id := uint8(i)
		_, ok := server.clients[id]

		if !ok {
			return id, nil
		}
	}

	return uint8(0), errors.New("Client limit reached!")
}

func CreateServer() Server {
	return Server{
		clients: make(map[uint8]*Client),
		in:      make(chan ClientMessage),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			Subprotocols:    []string{"websocket"},
			Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
				log.Printf("Error: %v %v", status, reason)
			},
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		encoder: codec.CreateBinaryEncoder([]codec.RegisteredCodec{
			{0, "me:id", &codec.Int8Codec{}},
			{0, "me:name", &codec.StringCodec{}},
			{0, "me:position", codec.CreatePositionCodec()},
			{0, "client:add", codec.CreateClientAddCodec()},
			{0, "client:remove", &codec.Int8Codec{}},
			{0, "client:name", codec.CreateClientNameCodec()},
			{0, "client:position", codec.CreateClientPositionCodec()},
			{0, "say", &codec.StringCodec{}},
		}, &codec.Int8Codec{}),
	}
}
