package websocket

import (
	"log"
	"cloud.google.com/go/pubsub"
)

// All the events received from the UI will be read from here 
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	// Inifite loop to constantly keep an eye on events received from UI for a particular user
	for {
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Println("Error while reading messages for connection ID ", c.ID, "Error  -> " ,  err)
			return
		}

		var data pubsub.Message
		var event string = string(p)
		log.Println("👉🏻 Event received from the UI 💻 ->", event)
		/*
		 * Events came from UI can be:
		 * 1. createWorkflow
		 * 2. error
		 * 3. restart
		 * 4. end
		 */
		data = createPubSubMsg(c.ID, event)
		publishToTLService(data, event)
	}
}
