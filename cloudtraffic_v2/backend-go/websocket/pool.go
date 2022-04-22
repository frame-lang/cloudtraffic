package websocket

import (
	"fmt"
	"log"
	"os"

	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/trafficlight"
)

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
			fmt.Println("Client registered, Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Connection.WriteJSON(Response{
					Type:           "addedInPool",
					Message:        client.ID,
					ConnectedUsers: len(pool.Clients),
				})
			}
			break
		case client := <-pool.Unregister:
			fmt.Println("Client unregistered, Size of Connection Pool: ", len(pool.Clients))
			for singleClient, _ := range pool.Clients {
				if client.ID == singleClient.ID {
					trafficlight.TrafficLights[singleClient.ID].Stop()
					fileName := trafficlight.GetFileName(client.ID)
					err := os.Remove(fileName)
					if err != nil {
						log.Fatal(err)
					}
					delete(pool.Clients, client)
				}
				client.Connection.WriteJSON(Response{
					Type:           "removedFromPool",
					Message:        singleClient.ID,
					ConnectedUsers: len(pool.Clients),
				})
			}
			break

			// No Broadcast functionality use currently

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
