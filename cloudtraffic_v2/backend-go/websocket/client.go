package websocket

import (
	"log"
	"cloud.google.com/go/pubsub"
)


func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	for {
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var data pubsub.Message
		
		if string(p) == "init" {
			data = createPubSubMsg(c.ID, "init")
		} else if string(p) == "error" {
			data = createPubSubMsg(c.ID, "error")
		} else if string(p) == "restart" {
			data = createPubSubMsg(c.ID, "restart")
		} else if string(p) == "end" {
			data = createPubSubMsg(c.ID, "end")
		}

		publishToTLService(data)
	}
}
