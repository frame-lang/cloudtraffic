// Package p contains a Pub/Sub Cloud Function.
package trafficlight

import (
	"context"
	"log"
	"strconv"
	"time"
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
	connectionID string
    redisPool *redis.Pool
	cloudFunctionID string
)

func init() {
	// For checking when a new intance is created
	cloudFunctionID = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	log.Println("A new cloud function is being inilialized: ", cloudFunctionID)

	// Initialize Redis
	var redisError error
	redisPool, redisError = initializeRedis()
	if redisError != nil {
		log.Printf("initializeRedis: %v", redisError)
		return
	}
}

func EntryPoint(_ctx context.Context, m PubSubMessage) error {
	connectionID = m.Attributes.ConnectionID
	var event string = m.Attributes.Event
	var createWorkflow bool = false
	log.Println("Connection ID ->", connectionID, ", Event ->", event, "Cloud function ID ->", cloudFunctionID)

	if event == "createWorkflow" {
		createWorkflow = true
	}

	trafficLightManager := NewTrafficLightManager(createWorkflow)

	if event == "tick" {
		trafficLightManager.Tick()
	} else if event == "error" {
		trafficLightManager.SystemError()
	} else if event == "restart" {
		trafficLightManager.SystemRestart()
	} else if event == "end" {
		trafficLightManager.End()
	} else if event == "connectionClosed" {
		trafficLightManager.ConnectionClosed()
	}
	
	return nil
}
