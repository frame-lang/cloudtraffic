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
		
		/*
		 * Events came from UI (string(p)) can be:
		 * 1. createWorkflow
		 * 2. error
		 * 3. restart
		 * 4. end
		 */
		data = createPubSubMsg(c.ID, string(p))
		publishToTLService(data)
	}
}
