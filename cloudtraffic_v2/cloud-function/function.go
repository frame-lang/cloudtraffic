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

	trafficLightMom := NewTrafficLightMom()
	var userID string = m.Attributes.ClientID
	var event string = m.Attributes.Event

	log.Println("Client ID ->", userID, ", Event ->", event)
	setUserID(userID)

	if event == "init" {
		trafficLightMom.Init()
	} else if event == "tick" {
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
