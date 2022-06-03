// A pool of users which are active on the web-socket (All connected users) 
package websocket

import (
	"log"
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
			log.Println("New user registered, Size of connection pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Connection.WriteJSON(PoolResponse{
					Type:           "addedInPool",
					NewUserID:        client.ID,
					TotalConnectedUsers: len(pool.Clients),
				})
			}
			break
		case client := <-pool.Unregister:
			log.Println("User ", client.ID, " is unregistered, Size of connection pool: ", len(pool.Clients))
			for singleClient, _ := range pool.Clients {
				if client.ID == singleClient.ID {
					// Stop the timer for particular user
					var activeUser *Client = Clients[client.ID]
					activeUser.Stopper <- true

					// Send event to remove user data from Redis
					data := createPubSubMsg(client.ID, "connectionClosed")
					publishToTLService(data, "connectionClosed")

					// Remove user from the active user's pool
					delete(pool.Clients, client)
				}
				client.Connection.WriteJSON(PoolResponse{
					Type:           "removedFromPool",
					NewUserID:        singleClient.ID,
					TotalConnectedUsers: len(pool.Clients),
				})
			}
			break
		}
	}
}
