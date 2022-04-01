package trafficlight

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var TrafficLights = make(map[string]TrafficLightMom)

func CreateNewTrafficLight (clientId string, conn *websocket.Conn) {
	TrafficLights[clientId] = NewTrafficLightMom(clientId, conn)
	fmt.Println(TrafficLights)
}