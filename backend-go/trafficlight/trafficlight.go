package trafficlight

import (
	"log"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

var TrafficLights = make(map[string]TrafficLightMom)

func CreateNewTrafficLight (clientId string, conn *websocket.Conn) {
	TrafficLights[clientId] = NewTrafficLightMom(clientId, conn)
}

type StateResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Loading bool   `json:"loading"`
}

func createResponse(state string, message string, loading bool) StateResponse {
	return StateResponse {
		Name:    state,
		Message: message,
		Loading: loading,
	}
}

func sendResponse(data StateResponse, conn *websocket.Conn) {
	if err := conn.WriteJSON(data); err != nil {
		log.Println(err)
		return
	}
}

func setInterval(p interface{}, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				log.Println("stop setInterval")
				return
			case <-ticker.C:
				reflect.ValueOf(p).Call([]reflect.Value{})
			}
		}

	}()

	// return the bool channel to use it as a stopper
	return stopIt
}