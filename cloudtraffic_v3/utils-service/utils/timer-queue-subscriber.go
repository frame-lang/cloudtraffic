package utils

import (
	"fmt"
	"encoding/json"
	"github.com/streadway/amqp"
)

type TimerData struct {
	ConnectionID	string
	Event			string 
}

func CreateTimerQueueSubscriber(channel *amqp.Channel) {
	_, err := channel.QueueDeclare(
		TIMER_QUEUE,
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
		TIMER_QUEUE,
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
			fmt.Printf("💁🏻 Received Message (Timer queue): %s\n", d.Body)

			var data TimerData
			err = json.Unmarshal(d.Body, &data)

			var activeUser *Client = Clients[data.ConnectionID]
			if activeUser == nil {
				return
			}

			if data.Event == "startWorkingTimer" {
				activeUser.Timer = setInterval(tick, activeUser.WorkingTimer, data.ConnectionID)
				activeUser.TickInProgress = true
			} else if data.Event == "startFlashingTimer" {
				activeUser.Timer = setInterval(tick, activeUser.FlashingTimer, data.ConnectionID)
				activeUser.TickInProgress = true
			} else if data.Event == "stopWorkingTimer" || data.Event == "stopFlashingTimer" {
				activeUser.Timer <- true
				activeUser.TickInProgress = false
			}
		}
	}()

	fmt.Println("⏳ Waiting for messages (Timer Queue)...")
	<- forever
}