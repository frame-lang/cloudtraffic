package trafficlight

import (
	"fmt"
	"os"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/streadway/amqp"
	"encoding/json"
)

type UIResponseData struct {
	ConnectionID	string
	Event			string 
	Color			string
	Loading 		bool
}

type TimerData struct {
	ConnectionID	string
	Event			string 
}

func createPayloadForUIResponseQueue(event string, color string, loading bool) []byte {
	var data = UIResponseData {connectionID, event, color, loading}
	byteData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return byteData
}

func createPayloadForTimerQueue(event string) []byte {
	var data = TimerData {connectionID, event}
	byteData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return byteData
}

func getRaabitURL() string {
	host := os.Getenv("RABBIT_HOST")
	port := os.Getenv("RABBIT_PORT")
	userName := os.Getenv("RABBIT_USERNAME")
	password := os.Getenv("RABBIT_PASSWORD")

	return "amqp://" + userName + ":" + password + "@" + host + ":" + port
}

func sendMessage(
	queueName	string,
	payload 	[]byte,
) {
	fmt.Println("queueName", queueName)
	fmt.Println("payload", payload)

	rabbitURL := getRaabitURL()
	conn, err := amqp.Dial(rabbitURL)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	
	fmt.Println("Successfully connected to RabbitMQ ✅...")

	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	queue, err := channel.QueueDeclare(
		queueName,
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
	fmt.Println("queue", queue)

	err = channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: payload,
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully published message to:", queueName)

	channel.Close()
	conn.Close()
}

func initializeRedis() (*redis.Pool, error) {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
			return nil, errors.New("REDISHOST must be set")
	}
	redisPort := os.Getenv("REDIS_PORT")
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