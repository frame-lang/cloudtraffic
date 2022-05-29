// Package p contains a Pub/Sub Cloud Function.
package trafficlight

import (
	"context"
	"log"
	"os"

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

func EntryPoint(ctx context.Context, m PubSubMessage) error {
	connectionID = m.Attributes.ConnectionID
	var event string = m.Attributes.Event
	var isInit bool = false
	log.Println("Connection ID ->", connectionID, ", Event ->", event)

	if event == "init" {
		isInit = true
	}

	trafficLightMom := NewTrafficLightMom(isInit)

	if event == "tick" {
		trafficLightMom.Tick()
	} else if event == "error" {
		trafficLightMom.SystemError()
	} else if event == "restart" {
		trafficLightMom.SystemRestart()
	} else if event == "end" {
		trafficLightMom.End()
	}
	
	return nil
}
