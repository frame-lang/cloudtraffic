package trafficlight

import (
	"time"
	"log"

	"github.com/gomodule/redigo/redis"
)

func (m *trafficLightStruct) initTrafficLight() {
	m.mom.InitTrafficLight()
}

func (m *trafficLightStruct) destroyTrafficLight() {
	m.mom.DestroyTrafficLight()
}

func (m *trafficLightStruct) enterRed() {
	m.ChangeColor("red")
	m.mom.EnterRed()
}

func (m *trafficLightStruct) enterGreen() {
	m.ChangeColor("green")
	m.mom.EnterGreen()
}

func (m *trafficLightStruct) enterYellow() {
	m.ChangeColor("yellow")
	m.mom.EnterYellow()
}

func (m *trafficLightStruct) enterFlashingRed() {
	m.ChangeColor("red")
	m.mom.EnterFlashingRed()
}

func (m *trafficLightStruct) startWorkingTimer() {
	m.mom.StartWorkingTimer()
}

func (m *trafficLightStruct) stopWorkingTimer() {
	m.mom.StopWorkingTimer()
}

func (m *trafficLightStruct) startFlashingTimer() {
	m.mom.StartFlashingTimer()
}
func (m *trafficLightStruct) stopFlashingTimer() {
	m.mom.StopFlashingTimer()
}

func (m *trafficLightStruct) changeFlashingAnimation() {
	flashColor := ""
	if m.flashColor == "red" {
		flashColor = "default"
	} else {
		flashColor = "red"
	}

	m.ChangeColor(flashColor)
	m.mom.ChangeFlashingAnimation()
}

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

	data, err := redis.String(conn.Do("GET", userID))
	if err != nil {
		log.Println(err)
	}
	log.Println("Data Received from Redis for User ", userID, "->", data)

	return []byte(data)
}

func (m *trafficLightMomStruct) setInRedis(data []byte) {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", userID, string(data))
	if err != nil {
			log.Printf("redis.Int: %v", err)
	}
	log.Println("Worflow saved to Redis for User ID", userID, ".")
}