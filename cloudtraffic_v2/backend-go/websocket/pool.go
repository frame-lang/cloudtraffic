package websocket

import (
	"fmt"
)

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
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
					// *** Event to clear data of user and stop timer
					fmt.Println("Unregistered User ID ->", client.ID)
					Users[client.ID].Stopper <- true
					delete(pool.Clients, client)
				}
				client.Connection.WriteJSON(Response{
					Type:           "removedFromPool",
					Message:        singleClient.ID,
					ConnectedUsers: len(pool.Clients),
				})
			}
			break
		}
	}
}
