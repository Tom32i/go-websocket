package server

import (
    // "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    // "time"
)

type Server struct {
    upgrader websocket.Upgrader
    id uint8
    clients map[uint8]Client
    in chan Message
    out chan []byte
    encoder BinaryEncoder
}

func (server *Server) Handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("upgrading connexion")
    socket, err := server.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    //defer c.Close(websocket.StatusInternalError, "the sky is falling")

    client := server.createClient(socket)
    client.write(server.encoder.encode("id", client.id))

    go client.run(server)
}

func (server *Server) Run() {
    log.Printf("Server running")
    for {
        select {
            case m := <-server.in:
                log.Printf("message in: %v", m)
                switch m.Name {
                    case "name":
                        m.client.setName(m.Data.(string))
                        log.Printf("client name: %v", m.client.name)
                        /*data := map[string]interface{}{
                            "id": m.client.id,
                            "name": m.data,
                        }
                        message := Message{
                            client: m.client,
                            name: "name",
                            data: data,
                        }
                        server.out <- message*/
                }
            case m := <-server.out:
                log.Printf("message out: %v", m)
        }
    }
}

func (server *Server) createClient(socket *websocket.Conn) Client {
    server.id += 1

    c := Client{
        id: server.id,
        socket: socket,
        encoder: server.encoder,
        //name: "Tom32i",
    }

    server.clients[c.id] = c

    log.Printf("Client #%d joined.", c.id)

    return c
}

func (server *Server) removeClient(c *Client) {
    delete(server.clients, c.id)
    log.Printf("Client #%d left.", c.id)
}

func (server Server) writeAll(name string, data interface{}) {
    buffer := server.encoder.encode(name, data)

    for _, c := range server.clients {
        c.write(buffer)
    }
}

func CreateServer() Server {
    return Server{
        id: 0,
        clients: make(map[uint8]Client),
        in: make(chan Message, 16),
        out: make(chan []byte, 16),
        upgrader: websocket.Upgrader{
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
        },
        encoder: createBinaryEncoder([]Codec{
            Int8Codec{BaseCodec{0, "id"}},
            StringCodec{BaseCodec{1, "name"}},
            StringCodec{BaseCodec{2, "say"}},
        }, Int8Codec{}),
    }
}

