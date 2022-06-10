package websocket

import (
	"context"
	"log"
	"strconv"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)

// Use to pull the events emmited from Cloud function (TL service) 
func PullMsgs() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Println("Error while creating connection to Cloud PubSub ->", err)
	}
	defer client.Close()

	sub := client.Subscription(SUBSCRIPTION_ID)
	var res StateResponse

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		log.Println("👉🏻 Event received from TL service ☁️ ->", string(msg.Data), "\n")
		connectionID :=  msg.Attributes["ConnectionID"]
		var activeUser *Client = Clients[connectionID]
		if activeUser == nil {
			return
		}
		
		if string(msg.Data) == "timerEvent" {
			event := msg.Attributes["Event"]
			log.Println("🕙 Timer Event received for connection ID", connectionID, "->", event)
			if event == "startWorkingTimer" {
				activeUser.Stopper = setInterval(tick, 3*time.Second, connectionID)
			} else if event == "startFlashingTimer" {
				activeUser.Stopper = setInterval(tick, 2*time.Second, connectionID)
			} else if event == "stopWorkingTimer" || event == "stopFlashingTimer" {
				activeUser.Stopper <- true
			}
			return
		}

		loading, err := strconv.ParseBool(msg.Attributes["Loading"])
        if err != nil {
            log.Fatal(err)
        }
		res = StateResponse {
			Name: msg.Attributes["Event"],
			Color: msg.Attributes["Color"],
			Loading: loading,
		}

		if err := activeUser.Connection.WriteJSON(res); err != nil {
			log.Println(err)
			return
		}
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})

	if err != nil {
		log.Println("Error while receiving events", err)
	}
	log.Println("Received ", received, " messages")
	return
}
