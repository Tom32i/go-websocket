package server

type Message struct {
    client *Client
    Name string
    Data interface{}
}

// type ClientNameMessage struct {
//     Message
//     data interface{
//         name string `json:"name"`
//         id string `json:"id"`
//     }
// }
