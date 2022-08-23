package utils

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/streadway/amqp"
)

type UIResponseData struct {
	ConnectionID	string
	Event			string 
	Color			string
	Loading 		bool
}

type UIPayload struct {
	Name    string `json:"name"`
	Color string `json:"color"`
	Loading bool   `json:"loading"`
}

func CreateUIResponseQueueSubscriber(channel *amqp.Channel) {
	_, err := channel.QueueDeclare(
		UI_RESPONSE_QUEUE,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	msgs, err := channel.Consume(
		UI_RESPONSE_QUEUE,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("💁🏻 Received Message (UI Response queue): %s\n", d.Body)

			var data UIResponseData
			err = json.Unmarshal(d.Body, &data)

			var activeUser *Client = Clients[data.ConnectionID]
			if activeUser == nil {
				return
			}

			ui_payload := UIPayload {
				Name: data.Event,
				Color: data.Color,
				Loading: data.Loading,
			}

			if err := activeUser.Connection.WriteJSON(ui_payload); err != nil {
				log.Println(err)
				return
			}
		}
	}()

	fmt.Println("⏳ Waiting for messages (UI Response Queue)...")
	<- forever
}