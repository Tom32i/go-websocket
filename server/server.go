package server

import (
    // "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    // "time"
    "curvygo/server/codec"
)

type Server struct {
    upgrader websocket.Upgrader
    id uint8
    clients map[uint8]*Client
    in chan ClientMessage
    encoder codec.BinaryEncoder
}

func (server *Server) Handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Upgrading connexion")
    socket, err := server.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    //defer c.Close(websocket.StatusInternalError, "the sky is falling")

    client := server.createClient(socket)
    server.init(client)

    go client.run(server)
}

func (server *Server) Run() {
    log.Printf("Server running")
    for {
        select {
            case m := <-server.in:
                log.Printf("message in: %v (%v)", m.message.Name, m.message.Data)
                switch m.message.Name {
                    case "me:name":
                        m.client.setName(m.message.Data.(string))
                        log.Printf("Client #%v name is '%v'.", m.client.id, m.client.name)
                        // server.writeAll("say", "Test € !")
                        server.writeAll(
                            "client:name",
                            codec.ClientNameMessage { m.client.id, m.client.name },
                        )
                }
        }
    }
}

func (server *Server) createClient(socket *websocket.Conn) *Client {
    server.id++

    c := Client{
        id: server.id,
        socket: socket,
        encoder: server.encoder,
    }

    server.clients[c.id] = &c

    log.Printf("Client #%d joined.", c.id)

    return &c
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
    server.writeAll("client:add", codec.ClientAddMessage { client.id, client.name })

    // Send the client its id
    client.write(server.encoder.Encode("me:id", client.id))

    // Send the clients the current client list
    for _, c := range server.clients {
        if c.id != client.id {
            client.write(server.encoder.Encode("client:add", codec.ClientAddMessage { c.id, c.name }))
        }
    }
}

func CreateServer() Server {
    return Server{
        id: 0,
        clients: make(map[uint8]*Client),
        in: make(chan ClientMessage, 16),
        upgrader: websocket.Upgrader{
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
            CheckOrigin: func(r *http.Request) bool {
                return true
            },
        },
        /*encoder: codec.CreateBinaryEncoder([]codec.Codec{
            codec.Int8Codec{codec.BaseCodec{0, "me:id"}},
            codec.StringCodec{codec.BaseCodec{1, "me:name"}},
            codec.createClientAddCodec(2, "client:add"),
            codec.Int8Codec{codec.BaseCodec{3, "client:remove"}},
            codec.StringCodec{codec.BaseCodec{4, "client:name"}},
            codec.StringCodec{codec.BaseCodec{5, "say"}},
        }, codec.Int8Codec{}),*/
        encoder: codec.CreateBinaryEncoder([]codec.RegisteredCodec{
            codec.RegisteredCodec{0, "me:id", codec.Int8Codec{}},
            codec.RegisteredCodec{0, "me:name", codec.StringCodec{}},
            codec.RegisteredCodec{0, "client:add", codec.CreateClientAddCodec()},
            codec.RegisteredCodec{0, "client:remove", codec.Int8Codec{}},
            codec.RegisteredCodec{0, "client:name", codec.CreateClientNameCodec()},
            codec.RegisteredCodec{0, "say", codec.StringCodec{}},
        }, codec.Int8Codec{}),
    }
}

