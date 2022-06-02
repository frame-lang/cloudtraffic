// Package p contains a Pub/Sub Cloud Function.
package trafficlight

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"
	"cloud.google.com/go/pubsub"
	"github.com/gomodule/redigo/redis"
)

type AttributeType struct {
	ConnectionID string `json:"connectionID"`
	Event string `json:"event"`
}

type PubSubMessage struct {
	Data       []byte            `json:"data"`
	Attributes AttributeType `json:"attributes"`
}

var (
	topic *pubsub.Topic
	client *pubsub.Client
	connectionID string
    redisPool *redis.Pool
	cloudFunctionID string
	ctx context.Context = context.Background()
)

func init() {
	// For testing
	cloudFunctionID = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	log.Println("A new cloud function is being inilialized: ", cloudFunctionID)

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

func EntryPoint(_ctx context.Context, m PubSubMessage) error {
	connectionID = m.Attributes.ConnectionID
	var event string = m.Attributes.Event
	var isNewWorkflow bool = false
	log.Println("Connection ID ->", connectionID, ", Event ->", event, "Cloud function ID ->", cloudFunctionID)

	if event == "init" {
		isNewWorkflow = true
	}

	trafficLightManager := NewTrafficLightManager(isNewWorkflow)

	if event == "tick" {
		trafficLightManager.Tick()
	} else if event == "error" {
		trafficLightManager.SystemError()
	} else if event == "restart" {
		trafficLightManager.SystemRestart()
	} else if event == "end" {
		trafficLightManager.End()
	}
	
	return nil
}
