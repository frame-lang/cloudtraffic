package trafficlight

import (
	"log"
	"github.com/gomodule/redigo/redis"
)

func (m *trafficLightManagerStruct) initTrafficLight() {
	sendMessage("sendResponseToUI", "begin", "", "true")
}

func (m *trafficLightManagerStruct) destroyTrafficLight() {
	sendMessage("sendResponseToUI", "end", "", "false")
}

func (m *trafficLightManagerStruct) enterRed() {
	color := m.trafficLight.GetColor()
	sendMessage("sendResponseToUI", "working", color, "false")
}

func (m *trafficLightManagerStruct) enterGreen() {
	color := m.trafficLight.GetColor()
	sendMessage("sendResponseToUI", "working", color, "false")
}

func (m *trafficLightManagerStruct) enterYellow() {
	color := m.trafficLight.GetColor()
	sendMessage("sendResponseToUI", "working", color, "false")
}

func (m *trafficLightManagerStruct) enterFlashingRed() {
	color := m.trafficLight.GetColor()
	sendMessage("sendResponseToUI", "error", color, "false")
}

func (m *trafficLightManagerStruct) startWorkingTimer() {
	sendMessage("timerEvent", "startWorkingTimer", "", "")
}

func (m *trafficLightManagerStruct) stopWorkingTimer() {
	sendMessage("timerEvent", "stopWorkingTimer", "", "")
}

func (m *trafficLightManagerStruct) startFlashingTimer() {
	sendMessage("timerEvent", "startFlashingTimer", "", "")
}

func (m *trafficLightManagerStruct) stopFlashingTimer() {
	sendMessage("timerEvent", "stopFlashingTimer", "", "")
}

func (m *trafficLightManagerStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	sendMessage("sendResponseToUI", "error", color, "false")
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