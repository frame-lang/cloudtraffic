package websocket

import (
    "github.com/gorilla/websocket"
)

type PoolResponse struct {
    Type string    `json:"type"`
    NewUserID string `json:"newUserID"`
    TotalConnectedUsers int `json:"totalConnectedUsers"`
}

type Pool struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
}

type Client struct {
    ID   string
    Pool *Pool
	Connection *websocket.Conn
    Timer chan<- bool
    TickInProgress bool
    WorkingTimer string
    FlashingTimer string
}

type StateResponse struct {
	Name    string `json:"name"`
	Color string `json:"color"`
	Loading bool   `json:"loading"`
}

type ResponseMessage struct {
	Data       []byte            `json:"data"`
	Attributes StateResponse `json:"attributes"`
}

type UIMessage struct {
    Event	string `json:"event"`
    State 	string `json:"state"`
    Data 	string `json:"data"`
}