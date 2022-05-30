package trafficlight

import (
	"context"
	"fmt"
	"os"
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/gomodule/redigo/redis"
)

var ctx context.Context = context.Background()

func publishResponse(state string, message string, loading string) {
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("sendResponse"),
		Attributes: map[string]string {
			"ConnectionID": connectionID,
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
	fmt.Println("Published a message to Utils service; msg ID: ", id)
}

func publishTimerEvent(eventName string, timerType string) {
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(eventName),
		Attributes: map[string]string {
			"ConnectionID": connectionID,
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