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
		
		if string(msg.Data) == "enableTimer" {
			log.Println("🕙 Enable timer for connection ID ", connectionID)
			timerType := msg.Attributes["TimerType"]
			if timerType == "workingTimer" {
				activeUser.Stopper = setInterval(tick, 3*time.Second, connectionID)
			} else if timerType == "flashingTimer" {
				activeUser.Stopper = setInterval(tick, 2*time.Second, connectionID)
			}
			return
		}

		if string(msg.Data) == "disableTimer" {
			log.Println("🕙 Disable timer for connection ID ", connectionID)
			activeUser.Stopper <- true
			return
		}

		loading, err := strconv.ParseBool(msg.Attributes["Loading"])
        if err != nil {
            log.Fatal(err)
        }
		res = StateResponse {
			Name: msg.Attributes["Name"],
			Message: msg.Attributes["Message"],
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
