// Contains the helper functions used in the package
package utils

import (
	"log"
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"bytes"
)

type UIEventData struct {
    ConnectionID	string
    Event		string
}

var Clients = make(map[string]*Client)

func AddClient(ID string, client *Client) {
	Clients[ID] = client
}

func setInterval(p func(string), intervalInSec string, connectionID string) chan<- bool {
	interval, _ := time.ParseDuration(intervalInSec)
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				log.Println("Stopped Timer")
				return
			case <-ticker.C:
				p(connectionID)
			}
		}

	}()

	// return the bool channel to use it as a stopper
	return stopIt
}

func tick(connectionID string) {
	sendEventToTLService(connectionID, "tick")
}

func getRaabitURL() string {
	host := RABBIT_HOST
	port := RABBIT_PORT
	userName := RABBIT_USERNAME
	password := RABBIT_PASSWORD

	return "amqp://" + userName + ":" + password + "@" + host + ":" + port
}

// Trigger Cloud function
func sendEventToTLService(connectionID string, event string) {
	var data = UIEventData {connectionID, event}
	byteData, err := json.Marshal(data)

	resp, err := http.Post(CLOUD_FUNC_URL, "application/json", bytes.NewBuffer(byteData))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println("Response from TL Service:", res["json"])
}