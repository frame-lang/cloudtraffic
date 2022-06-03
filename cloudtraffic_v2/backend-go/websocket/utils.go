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

func setInterval(p func(string), interval time.Duration, connectionID string) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				log.Println("Stopped Timer")
				return
			case <-ticker.C:
				p(connectionID)
			}
		}

	}()

	// return the bool channel to use it as a stopper
	return stopIt
}

func tick(connectionID string) {
	log.Println("Tick called for connectionID ->", connectionID)
	data := createPubSubMsg(connectionID, "tick")
	publishToTLService(data)
}

func createPubSubMsg(connectionID string, eventName string) pubsub.Message {
	return pubsub.Message{
		Data: []byte("Data from Go Back-end service"),
		Attributes: map[string]string{
			"connectionID": connectionID,
			"event": eventName,
		},
	}
}