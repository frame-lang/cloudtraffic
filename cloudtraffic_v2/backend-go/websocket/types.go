package websocket

import (
    "github.com/gorilla/websocket"
)

type Response struct {
    Type string    `json:"type"`
    Message string `json:"message"`
    ConnectedUsers int `json:"connectedUsers"`
}

type Pool struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    // Broadcast  chan Request
}

type Client struct {
    ID   string
    Pool *Pool
	Connection *websocket.Conn
    Stopper chan<- bool
}