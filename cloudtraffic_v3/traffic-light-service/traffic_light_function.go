// Package p contains a Pub/Sub Cloud Function.
package trafficlight

import (
	"log"
	"strconv"
	"time"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
)

type AttributeType struct {
	ConnectionID string `json:"connectionID"`
	Event string `json:"event"`
}

var (
	connectionID string
    redisPool *redis.Pool
	cloudFunctionID string
)

func init() {
	// For checking when a new intance is created
	cloudFunctionID = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	log.Println("A new cloud function is being inilialized: ", cloudFunctionID)

	// Initialize Redis
	var redisError error
	redisPool, redisError = initializeRedis()
	if redisError != nil {
		log.Printf("initializeRedis: %v", redisError)
		return
	}
}


func EntryPoint(w http.ResponseWriter, r *http.Request) {
	var data struct {
		ConnectionID string `json:"connectionID"`
		Event string `json:"event"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "Hello World!")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	connectionID = data.ConnectionID
	var event string = data.Event
	var createWorkflow bool = false
	log.Println("Connection ID ->", connectionID, ", Event ->", event, "Cloud function ID ->", cloudFunctionID)

	if event == "createWorkflow" {
		createWorkflow = true
	}

	trafficLightManager := NewTrafficLightManager(createWorkflow)

	if event == "tick" {
		trafficLightManager.Tick()
	} else if event == "error" {
		trafficLightManager.SystemError()
	} else if event == "restart" {
		trafficLightManager.SystemRestart()
	} else if event == "end" {
		trafficLightManager.End()
	} else if event == "connectionClosed" {
		trafficLightManager.ConnectionClosed()
	}
	
	fmt.Fprint(w, html.EscapeString("Function Triggered!"))
}
