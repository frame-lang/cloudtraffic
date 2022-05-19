package trafficlight

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

type StateResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Loading string   `json:"loading"`
}

type ResponseMessage struct {
	Data       []byte            `json:"data"`
	Attributes StateResponse `json:"attributes"`
}

var (
	topic *pubsub.Topic
	client *pubsub.Client
	ctx context.Context = context.Background()
	userID string
)

func init() {
	// err is pre-declared to avoid shadowing client.
	var err error

	// client is initialized with context.Background() because it should
	// persist between function invocations.
	client, err = pubsub.NewClient(ctx, "cloud-traffic-347207")
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}

	topic = client.Topic("cloudtraffic-utils-service-topic")
}

func setUserID(ID string) {
	userID = ID
}

func publishResponse(state string, message string, loading string) {
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("sendResponse"),
		Attributes: map[string]string {
			"UserID": userID,
			"Name": state,
			"Message": message,
			"Loading":loading,
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Errorf("Get: %v", err)
	}
	fmt.Println("Published a message; msg ID: ", id)
}

func publishTimerEvent(eventName string) {
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(eventName),
		Attributes: map[string]string {
			"UserID": userID,
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Errorf("Get: %v", err)
	}
	fmt.Println("Published a message; msg ID: ", id)
}