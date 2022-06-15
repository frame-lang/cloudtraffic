package trafficlight

import (
	"fmt"
	"os"
	"errors"
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/api/option"

)

func sendMessage(
	msgType string,
	event string,
	color string,
	loading string,
) {
	var ctx context.Context = context.Background()
	var client, err = pubsub.NewClient(
		ctx,
		os.Getenv("PROJECTID"),
		option.WithEndpoint("us-central1-pubsub.googleapis.com:443"),
	)
	if err != nil {
		fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	// Enable order messaging
	var topic *pubsub.Topic = client.Topic(os.Getenv("TOPICID"))
	topic.EnableMessageOrdering = true

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msgType),
		OrderingKey: "tl-service", // To publish messages in oder
		Attributes: map[string]string {
			"ConnectionID": connectionID,
			"Event": event,
			"Color": color,
			"Loading":loading,
		},
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Errorf("Get: %v", err)
	}
	fmt.Println("Published a message to Utils service; msg ID: ", id)
}

func initializeRedis() (*redis.Pool, error) {
	redisHost := os.Getenv("REDISHOST")
	if redisHost == "" {
			return nil, errors.New("REDISHOST must be set")
	}
	redisPort := os.Getenv("REDISPORT")
	if redisPort == "" {
			return nil, errors.New("REDISPORT must be set")
	}
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	const maxConnections = 10
	return &redis.Pool{
			MaxIdle: maxConnections,
			Dial: func() (redis.Conn, error) {
					c, err := redis.Dial("tcp", redisAddr)
					if err != nil {
							return nil, fmt.Errorf("redis.Dial: %v", err)
					}
					return c, err
			},
	}, nil
}