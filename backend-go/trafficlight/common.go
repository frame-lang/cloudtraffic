package trafficlight

import (
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

type commandStruct struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Loading bool   `json:"loading"`
}

var Stopper chan<- bool
var SocketConn *websocket.Conn
var MOM TrafficLightMom

func CreateNewTrafficLight() {
	MOM = NewTrafficLightMom()
}

func createResponse(state string, message string, loading bool) string {
	command := &commandStruct{
		Name:    state,
		Message: message,
		Loading: loading,
	}

	json, _ := json.Marshal(command)
	return string(json)
}

func sendResponse(data string) {
	if err := SocketConn.WriteJSON(data); err != nil {
		log.Println(err)
		return
	}
}

func SetInterval(p interface{}, interval time.Duration) chan<- bool {
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
