package trafficlight

import (
	"log"
	"github.com/gomodule/redigo/redis"
)

func (m *trafficLightManagerStruct) initTrafficLight() {
	var payload = createPayloadForUIResponseQueue("begin", "", true)
	sendMessage("UIResponseQueue", payload)
}

func (m *trafficLightManagerStruct) destroyTrafficLight() {
	var payload = createPayloadForUIResponseQueue("end", "", false)
	sendMessage("UIResponseQueue", payload)
}

func (m *trafficLightManagerStruct) enterRed() {
	color := m.trafficLight.GetColor()
	var payload = createPayloadForUIResponseQueue("working", color, false)
	sendMessage("UIResponseQueue", payload)
}

func (m *trafficLightManagerStruct) enterGreen() {
	color := m.trafficLight.GetColor()
	payload := createPayloadForUIResponseQueue("working", color, false)
	sendMessage("UIResponseQueue", payload)
}

func (m *trafficLightManagerStruct) enterYellow() {
	color := m.trafficLight.GetColor()
	payload := createPayloadForUIResponseQueue("working", color, false)
	sendMessage("UIResponseQueue", payload)
}

func (m *trafficLightManagerStruct) enterFlashingRed() {
	color := m.trafficLight.GetColor()
	payload := createPayloadForUIResponseQueue("error", color, false)
	sendMessage("UIResponseQueue", payload)
}

func (m *trafficLightManagerStruct) startWorkingTimer() {
	payload := createPayloadForTimerQueue("startWorkingTimer")
	sendMessage("TimerQueue", payload)
}

func (m *trafficLightManagerStruct) stopWorkingTimer() {
	payload := createPayloadForTimerQueue("stopWorkingTimer")
	sendMessage("TimerQueue", payload)
}

func (m *trafficLightManagerStruct) startFlashingTimer() {
	payload := createPayloadForTimerQueue("startFlashingTimer")
	sendMessage("TimerQueue", payload)
}

func (m *trafficLightManagerStruct) stopFlashingTimer() {
	payload := createPayloadForTimerQueue("stopFlashingTimer")
	sendMessage("TimerQueue", payload)
}

func (m *trafficLightManagerStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	payload := createPayloadForUIResponseQueue("error", color, false)
	sendMessage("UIResponseQueue", payload)
}

func (m *trafficLightManagerStruct) getWorkflowFromRedis() []byte {
	conn := redisPool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", connectionID))
	if err != nil {
		log.Println(err)
	}
	log.Println("Data Received from Redis for User ", connectionID, "->", data)

	return []byte(data)
}

func (m *trafficLightManagerStruct) setWorkflowInRedis(data []byte) {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", connectionID, string(data))
	if err != nil {
		log.Printf("Error while saving TL Data in Redis: %v", err)
		return
	}
	log.Println("Worflow saved to Redis for User ID", connectionID, ".")
}

func (m *trafficLightManagerStruct) removeWorkflowFromRedis() {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", connectionID)
	if err != nil {
		log.Printf("Error while removing data from Redis: %v", err)
		return
	}
	log.Println("Removed successfully from Redis,  User ID ", connectionID, ".")
}