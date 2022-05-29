package trafficlight

import (
	"time"
	"log"

	"github.com/gomodule/redigo/redis"
)

func (m *trafficLightMomStruct) initTrafficLight() {
	publishResponse("begin", "", "true")
	time.Sleep(2 * time.Second)
}

func (m *trafficLightMomStruct) destroyTrafficLight() {
	publishResponse("end", "", "false")
}

func (m *trafficLightMomStruct) enterRed() {
	color := m.trafficLight.GetColor()
	publishResponse("working", color, "false")
}

func (m *trafficLightMomStruct) enterGreen() {
	color := m.trafficLight.GetColor()
	publishResponse("working", color, "false")
}

func (m *trafficLightMomStruct) enterYellow() {
	color := m.trafficLight.GetColor()
	publishResponse("working", color, "false")
}

func (m *trafficLightMomStruct) enterFlashingRed() {
	color := m.trafficLight.GetColor()
	publishResponse("error", color, "false")
}

func (m *trafficLightMomStruct) startWorkingTimer() {
	publishTimerEvent("enableTimer", "workingTimer")
}

func (m *trafficLightMomStruct) stopWorkingTimer() {
	publishTimerEvent("disableTimer", "workingTimer")
}

func (m *trafficLightMomStruct) startFlashingTimer() {
	publishTimerEvent("enableTimer", "flashingTimer")
}

func (m *trafficLightMomStruct) stopFlashingTimer() {
	publishTimerEvent("disableTimer", "flashingTimer")
}


func (m *trafficLightMomStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	publishResponse("error", color, "false")
}

func (m *trafficLightMomStruct) getFromRedis() []byte {
	conn := redisPool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", connectionID))
	if err != nil {
		log.Println(err)
	}
	log.Println("Data Received from Redis for User ", connectionID, "->", data)

	return []byte(data)
}

func (m *trafficLightMomStruct) setInRedis(data []byte) {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", connectionID, string(data))
	if err != nil {
			log.Printf("redis.Int: %v", err)
	}
	log.Println("Worflow saved to Redis for User ID", connectionID, ".")
}