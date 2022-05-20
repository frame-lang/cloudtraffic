// Package p contains a Pub/Sub Cloud Function.
package trafficlight

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type AttributeType struct {
	ClientID string `json:"clientID"`
	Event string `json:"event"`
}

type PubSubMessage struct {
	Data       []byte            `json:"data"`
	Attributes AttributeType `json:"attributes"`
}

func EntryPoint(ctx context.Context, m PubSubMessage) error {
	log.Println("Client ID ->", m.Attributes.ClientID)


	var err error
	redisPool, err = initializeRedis()
	if err != nil {
			log.Printf("initializeRedis: %v", err)
	}

	trafficLightMom := NewTrafficLightMom()
	setUserID(m.Attributes.ClientID)

	conn := redisPool.Get()
	defer conn.Close()

	counter, err := conn.Do("APPEND", "test", "value")
	if err != nil {
			log.Printf("redis.Int: %v", err)
	}

	log.Println("counter", counter)

	log.Println("trafficLightMom", trafficLightMom)
	var event string = m.Attributes.Event

	if event == "init" {
		trafficLightMom.Init()
	} else if event == "tick" {
		trafficLightMom.Tick()
	} else if event == "error" {
		trafficLightMom.SystemError()
	} else if event == "restart" {
		trafficLightMom.SystemRestart()
	} else if event == "end" {
		trafficLightMom.Stop()
	}
	
	return nil
}
