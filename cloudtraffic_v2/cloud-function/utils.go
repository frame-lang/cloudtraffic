package trafficlight

import (
	"context"
	"fmt"
	"log"
	"os"
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/gomodule/redigo/redis"
)

var (
	topic *pubsub.Topic
	client *pubsub.Client
	ctx context.Context = context.Background()
	userID string
    redisPool *redis.Pool
)

func init() {
	// err is pre-declared to avoid shadowing client.
	var err error

	// client is initialized with context.Background() because it should
	// persist between function invocations.
	client, err = pubsub.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}

	topic = client.Topic(os.Getenv("TOPICID"))

	// Initialize Redis
	var redisError error
	redisPool, redisError = initializeRedis()
	if redisError != nil {
			log.Printf("initializeRedis: %v", err)
			return
	}
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

func publishTimerEvent(eventName string, timerType string) {
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(eventName),
		Attributes: map[string]string {
			"UserID": userID,
			"TimerType": timerType,
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