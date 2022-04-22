package cloudpubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func Publish(projectID string, topicID string, data pubsub.Message) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	fmt.Println("client", client)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
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