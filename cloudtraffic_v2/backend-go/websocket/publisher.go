package websocket

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func publishToTLService(data pubsub.Message) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(TOPIC_ID)
	result := t.Publish(ctx, &data)
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
			return fmt.Errorf("Get: %v", err)
	}
	fmt.Println("Published a message; msg ID: ", id)
	return nil
}