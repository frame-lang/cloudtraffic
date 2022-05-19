package websocket

import (
	"context"
	"fmt"
	"log"
	"strconv"
	// "strings"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
	// "github.com/gorilla/websocket"
)

type StateResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Loading bool   `json:"loading"`
}

type ResponseMessage struct {
	Data       []byte            `json:"data"`
	Attributes StateResponse `json:"attributes"`
}

func PullMsgs(projectID, subID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)
	var res StateResponse

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Println("Got message: ", string(msg.Data))
		userID :=  msg.Attributes["UserID"]
		var activeUser *Client = Users[userID]

		if activeUser == nil {
			return
		}
		
		if string(msg.Data) == "enableTimer" {
			log.Println("enableTimer")
			activeUser.Stopper = setInterval(tick, 5*time.Second, userID)
			return
		}

		if string(msg.Data) == "disableTimer" {
			log.Println("disableTimer")
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

		log.Println("res", res)

		if err := activeUser.Connection.WriteJSON(res); err != nil {
			log.Println(err)
			return
		}
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		fmt.Errorf("sub.Receive: %v", err)
	}
	fmt.Println("Received ", received, " messages")
	fmt.Println("res", res)
	return
}
