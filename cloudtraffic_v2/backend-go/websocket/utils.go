package websocket

import (
	"log"
	"time"
	"cloud.google.com/go/pubsub"
)


var Users = make(map[string]*Client)

func AddUser(ID string, user *Client) {
	Users[ID] = user
	log.Println("Users", Users)
}

func setInterval(p func(string), interval time.Duration, userID string) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				log.Println("Stopped Timer")
				return
			case <-ticker.C:
				p(userID)
			}
		}

	}()

	// return the bool channel to use it as a stopper
	return stopIt
}

func tick(userID string) {
	log.Println("Tick.....", userID)
	data := createPubSubMsg(userID, "tick")
	publishToTLService(data)
}

func createPubSubMsg(userId string, eventName string) pubsub.Message {
	return pubsub.Message{
		Data: []byte("Data from Go Back-end service"),
		Attributes: map[string]string{
			"clientID":userId,
			"event": eventName,
		},
	}
}