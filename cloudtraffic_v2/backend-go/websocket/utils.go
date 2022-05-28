package websocket

import (
	"log"
	"time"
	"cloud.google.com/go/pubsub"
)


var Clients = make(map[string]*Client)

func AddClient(ID string, client *Client) {
	Clients[ID] = client
}

func setInterval(p func(string), interval time.Duration, clientID string) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				log.Println("Stopped Timer")
				return
			case <-ticker.C:
				p(clientID)
			}
		}

	}()

	// return the bool channel to use it as a stopper
	return stopIt
}

func tick(clientID string) {
	log.Println("Tick.....", clientID)
	data := createPubSubMsg(clientID, "tick")
	publishToTLService(data)
}

func createPubSubMsg(clientID string, eventName string) pubsub.Message {
	return pubsub.Message{
		Data: []byte("Data from Go Back-end service"),
		Attributes: map[string]string{
			"clientID": clientID,
			"event": eventName,
		},
	}
}