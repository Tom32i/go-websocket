package server

type Message struct {
    client *Client
    Name string `json:"name"`
    Data interface{} `json:"data"`
}

// type ClientNameMessage struct {
//     Message
//     data interface{
//         name string `json:"name"`
//         id string `json:"id"`
//     }
// }
