package server

type Message struct {
    client *Client
    name string
    data any
}

// type ClientNameMessage struct {
//     Message
//     data interface{
//         name string `json:"name"`
//         id string `json:"id"`
//     }
// }
