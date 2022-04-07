package trafficlight

import (
	"log"
	"os"
	"path/filepath"
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

func RemoveContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}

func CreateDataDirIfNotExists () {
	if _, err := os.Stat(DataDirPath()); os.IsNotExist(err) {
		err := os.Mkdir(DataDirPath(), 0755)

		if (err != nil) {
			log.Fatal("Unable tocreate Data directory")
		}
	}
}

func DataDirPath () string {
	return filepath.Join("../", "data")
}

func GetFileName (name string) string {
	return filepath.Join("../", "data", name + ".json")
}