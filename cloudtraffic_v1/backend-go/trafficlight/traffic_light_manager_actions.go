package trafficlight

import (
	"log"
	"os"
	"time"
)

func (m *trafficLightMomStruct) initTrafficLight() {
	res := createResponse("begin", "", true)
	sendResponse(res, m.connection)
	time.Sleep(1 * time.Second)
}

func (m *trafficLightMomStruct) destroyTrafficLight() {
	res := createResponse("end", "", false)
	sendResponse(res, m.connection)
}

func (m *trafficLightMomStruct) saveInDisk(data []byte)  {
	fileName := GetFileName(m.clientId)
	jsonFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(data)
	jsonFile.Close()
}


func (m *trafficLightMomStruct) enterRed() {
	color := m.trafficLight.GetColor()
	res := createResponse("working", color, false)
	sendResponse(res, m.connection)
}

func (m *trafficLightMomStruct) enterGreen() {
	color := m.trafficLight.GetColor()
	res := createResponse("working", color, false)
	sendResponse(res, m.connection)
}

func (m *trafficLightMomStruct) enterYellow() {
	color := m.trafficLight.GetColor()
	res := createResponse("working", color, false)
	sendResponse(res, m.connection)
}

func (m *trafficLightMomStruct) enterFlashingRed() {
	color := m.trafficLight.GetColor()
	res := createResponse("error", color, false)
	sendResponse(res, m.connection)
}

func (m *trafficLightMomStruct) startWorkingTimer() {
	mom:= TrafficLights[m.clientId]
	m.stopper = setInterval(mom.Tick, 2*time.Second)
}

func (m *trafficLightMomStruct) stopWorkingTimer() {
	m.stopper <- true
}

func (m *trafficLightMomStruct) startFlashingTimer() {
	mom:= TrafficLights[m.clientId]
	m.stopper = setInterval(mom.Tick, 1*time.Second)
}

func (m *trafficLightMomStruct) stopFlashingTimer() {
	m.stopper <- true
}


func (m *trafficLightMomStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	res := createResponse("error", color, false)
	sendResponse(res, m.connection)
}

func (m *trafficLightMomStruct) loadFromDisk(clientId string) []byte {
	fileName := GetFileName(clientId)
	savedData, err1 := os.ReadFile(fileName)
	if err1 != nil {
		panic(err1)
	}

	return savedData
}