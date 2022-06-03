package websocket

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

// Publish events to Cloud function (Traffic Light service) via Cloud PubSub
func publishToTLService(data pubsub.Message, eventName string) {
	log.Println("👉🏻 Publishing to TL service ->", eventName)
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Println("Error while creating connection to Cloud PubSub ->", err)
	}
	defer client.Close()

	t := client.Topic(TOPIC_ID)
	result := t.Publish(ctx, &data)
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	_, err = result.Get(ctx)
	if err != nil {
			log.Println("Error while publishing data to Cloud PubSub ->", err)
	}
}