package trafficlight

func (m *trafficLightStruct) initTrafficLight() {
	m._manager_.InitTrafficLight()
}

func (m *trafficLightStruct) destroyTrafficLight() {
	m._manager_.DestroyTrafficLight()
}

func (m *trafficLightStruct) enterRed() {
	m.ChangeColor("red")
	m._manager_.EnterRed()
}

func (m *trafficLightStruct) enterGreen() {
	m.ChangeColor("green")
	m._manager_.EnterGreen()
}

func (m *trafficLightStruct) enterYellow() {
	m.ChangeColor("yellow")
	m._manager_.EnterYellow()
}

func (m *trafficLightStruct) enterFlashingRed() {
	m.ChangeColor("red")
	m._manager_.EnterFlashingRed()
}

func (m *trafficLightStruct) startWorkingTimer() {
	m._manager_.StartWorkingTimer()
}

func (m *trafficLightStruct) stopWorkingTimer() {
	m._manager_.StopWorkingTimer()
}

func (m *trafficLightStruct) startFlashingTimer() {
	m._manager_.StartFlashingTimer()
}
func (m *trafficLightStruct) stopFlashingTimer() {
	m._manager_.StopFlashingTimer()
}

func (m *trafficLightStruct) changeFlashingAnimation() {
	flashColor := ""
	if m.flashColor == "red" {
		flashColor = "default"
	} else {
		flashColor = "red"
	}

	m.ChangeColor(flashColor)
	m._manager_.ChangeFlashingAnimation()
}
