package server

type Message struct {
    client *Client
    Name string
    Data any
}

// type ClientNameMessage struct {
//     Message
//     data interface{
//         name string `json:"name"`
//         id string `json:"id"`
//     }
// }
