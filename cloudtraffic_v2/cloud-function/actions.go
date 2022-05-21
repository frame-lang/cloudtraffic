package trafficlight

import (
	"fmt"
	"time"
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
	fmt.Println("destroyTrafficLight")
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
	publishTimerEvent("enableTimer")
}

func (m *trafficLightMomStruct) stopWorkingTimer() {
	fmt.Println("stopWorkingTimer disableTimer")
	publishTimerEvent("disableTimer")
}

func (m *trafficLightMomStruct) startFlashingTimer() {
	publishTimerEvent("enableTimer")
}

func (m *trafficLightMomStruct) stopFlashingTimer() {
	publishTimerEvent("disableTimer")
}


func (m *trafficLightMomStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	publishResponse("error", color, "false")
}

func (m *trafficLightMomStruct) saveInDisk(data []byte)  {
	fmt.Println("saveInDisk")
	setInRedis(string(data))
}

func (m *trafficLightMomStruct) loadFromDisk() []byte {
	fmt.Println("loadFromDisk")
	var data []byte = getFromRedis()
	return data
}