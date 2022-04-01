package websocket

import "fmt"

func NewPool() *Pool {
    return &Pool{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        // Broadcast:  make(chan Request),
    }
}

func (pool *Pool) Start() {
    for {
        select {
            case client := <-pool.Register:
                pool.Clients[client] = true
                fmt.Println("Size of Connection Pool: ", len(pool.Clients))
                for client, _ := range pool.Clients {
                    fmt.Println("Client ->", client)
                    client.Conn.WriteJSON(Response{
                        Type: "addedInPool",
                        Message: client.ID,
                        ConnectedUsers: len(pool.Clients),
                    })
                }
                break
            case client := <-pool.Unregister:
                delete(pool.Clients, client)
                fmt.Println("Size of Connection Pool: ", len(pool.Clients))
                for client, _ := range pool.Clients {
                    client.Conn.WriteJSON(Response{
                        Type: "removedFromPool",
                        Message: client.ID,
                        ConnectedUsers: len(pool.Clients),
                    })
                }
                break
            // case message := <-pool.Broadcast:
            //     fmt.Println("Sending message to all clients in Pool")
            //     for client, _ := range pool.Clients {
            //         if err := client.Conn.WriteJSON(message); err != nil {
            //             fmt.Println(err)
            //             return
            //         }
            //     }
        }
    }
}