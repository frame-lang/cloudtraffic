package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitInstance struct {
	Conn *amqp.Connection
    Channel *amqp.Channel
}

func ConnectToRabbitMQ() RabbitInstance {
	rabbitURL := getRaabitURL()
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var rabbitConnection RabbitInstance
	rabbitConnection.Conn = conn
	
	fmt.Println("✅ Successfully connected to RabbitMQ...")

	channel, err := rabbitConnection.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	rabbitConnection.Channel = channel

	return rabbitConnection
}