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
	time.Sleep(3 * time.Second)
}

func (m *trafficLightMomStruct) destroyTrafficLight() {
	publishResponse("end", "", "false")
}

func (m *trafficLightMomStruct) saveInDisk(data []byte)  {
	fmt.Println("saveInDisk")

	// fileName := GetFileName(m.clientId)
	// jsonFile, err := os.Create(fileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer jsonFile.Close()

	// jsonFile.Write(data)
	// jsonFile.Close()
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
	publishTimerEvent("disableTimer")
}

func (m *trafficLightMomStruct) startFlashingTimer() {
	// mom:= TrafficLights[m.clientId]
	// m.stopper = setInterval(mom.Tick, 1*time.Second)
	// Call to util server to start timer
}

func (m *trafficLightMomStruct) stopFlashingTimer() {
	// m.stopper <- true
	// Call to util server to stop timer
}


func (m *trafficLightMomStruct) changeFlashingAnimation() {
	color := m.trafficLight.GetColor()
	publishResponse("error", color, "false")
}

func (m *trafficLightMomStruct) loadFromDisk(clientId string) []byte {
	fmt.Println("loadFromDisk")
	// fileName := GetFileName(clientId)
	// savedData, err1 := os.ReadFile(fileName)
	// if err1 != nil {
	// 	panic(err1)
	// }
	return []byte("testing")
	// return savedData
}