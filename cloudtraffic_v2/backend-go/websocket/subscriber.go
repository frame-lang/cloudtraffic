package websocket

import (
	"context"
	"log"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)

// Use to pull the events emmited from Cloud function (TL service) 
func PullMsgs() {
	log.Println("Pull subscriber is listening... 🎧")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Println("Error while creating connection to Cloud PubSub ->", err)
	}
	defer client.Close()

	sub := client.Subscription(SUBSCRIPTION_ID)
	var res StateResponse

	ok, err := sub.Exists(ctx)
	if err != nil {
		fmt.Errorf("Error while checking subscription exists", err)
	}

	// Create a new subscription, if not exsited
	if !ok {
		log.Println("Subscription not exists, creating a new one...")
		topic := client.Topic(UTILS_TOPIC_ID)

		sub, err := client.CreateSubscription(ctx, SUBSCRIPTION_ID, pubsub.SubscriptionConfig{
			Topic:                 topic,
			AckDeadline:           20 * time.Second,
			EnableMessageOrdering: true,
		})
		if err != nil {
				fmt.Errorf("Error while creating new subscription  %v ->", err)
		}
		log.Println("✅ Subscription created successfully ->", sub)
	}

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		atomic.AddInt32(&received, 1)
		msg.Ack()
		connectionID :=  msg.Attributes["ConnectionID"]
		var activeUser *Client = Clients[connectionID]
		if activeUser == nil {
			return
		}
		event := msg.Attributes["Event"]
		log.Println("👉🏻 Event received from TL service ☁️ ->", string(msg.Data))
		
		if string(msg.Data) == "timerEvent" {
			log.Println("🕙 Timer Event received for connection ID", connectionID, "->", event)
			if event == "startWorkingTimer" {
				activeUser.Timer = setInterval(tick, activeUser.WorkingTimer, connectionID)
				activeUser.TickInProgress = true
				} else if event == "startFlashingTimer" {
				activeUser.Timer = setInterval(tick, activeUser.FlashingTimer, connectionID)
				activeUser.TickInProgress = true
				} else if event == "stopWorkingTimer" || event == "stopFlashingTimer" {
				activeUser.Timer <- true
				activeUser.TickInProgress = false
			}
			return
		}

		loading, err := strconv.ParseBool(msg.Attributes["Loading"])
        if err != nil {
            log.Fatal(err)
        }
		res = StateResponse {
			Name: event,
			Color: msg.Attributes["Color"],
			Loading: loading,
		}

		if err := activeUser.Connection.WriteJSON(res); err != nil {
			log.Println(err)
			return
		}
	})

	if err != nil {
		log.Println("Error while receiving events", err)
	}
	log.Println("Received ", received, " messages")
	return
}