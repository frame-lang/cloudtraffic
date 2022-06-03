package trafficlight

import (
	"time"
	"log"

	"github.com/gomodule/redigo/redis"
)

func (m *trafficLightManagerStruct) initTrafficLight() {
	publishResponse("begin", "", "true")
	time.Sleep(2 * time.Second)
}

func (m *trafficLightManagerStruct) destroyTrafficLight() {
	publishResponse("end", "", "false")
}

func (m *trafficLightManagerStruct) enterRed() {
	color := m.trafficLight.GetColor()
	publishResponse("working", color, "false")
}

func (m *trafficLightManagerStruct) enterGreen() {
	color := m.trafficLight.GetColor()
	publishResponse("working", color, "false")
}

func (m *trafficLightManagerStruct) enterYellow() {
	color := m.trafficLight.GetColor()
	publishResponse("working", color, "false")
}

func (m *trafficLightManagerStruct) enterFlashingRed() {
	color := m.trafficLight.GetColor()
	publishResponse("error", color, "false")
}

func (m *trafficLightManagerStruct) startWorkingTimer() {
	publishTimerEvent("enableTimer", "workingTimer")
}

func (m *trafficLightManagerStruct) stopWorkingTimer() {
	publishTimerEvent("disableTimer", "workingTimer")
}

func (m *trafficLightManagerStruct) startFlashingTimer() {
	publishTimerEvent("enableTimer", "flashingTimer")
}

func (m *trafficLightManagerStruct) stopFlashingTimer() {
	publishTimerEvent("disableTimer", "flashingTimer")
}


func (m *trafficLightManagerStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	publishResponse("error", color, "false")
}

func (m *trafficLightManagerStruct) getFromRedis() []byte {
	conn := redisPool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", connectionID))
	if err != nil {
		log.Println(err)
	}
	log.Println("Data Received from Redis for User ", connectionID, "->", data)

	return []byte(data)
}

func (m *trafficLightManagerStruct) setInRedis(data []byte) {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", connectionID, string(data))
	if err != nil {
		log.Printf("Error while saving TL Data in Redis: %v", err)
		return
	}
	log.Println("Worflow saved to Redis for User ID", connectionID, ".")
}

func (m *trafficLightManagerStruct) removeFromRedis() {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", connectionID)
	if err != nil {
		log.Printf("Error while removing data from Redis: %v", err)
		return
	}
	log.Println("Removed successfully from Redis,  User ID ", connectionID, ".")
}