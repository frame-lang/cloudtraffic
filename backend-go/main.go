package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/trafficlight"
	"github.com/frame-lang/frame-demos/persistenttrafficlight/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	client := &websocket.Client{
		ID: timestamp,
		Pool: pool,
	}

	trafficlight.CreateNewTrafficLight(timestamp, conn)
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Traffic Light App v0.1.0")

	// Create a directory to store saved data (only if not exists currently)
	trafficlight.CreateDataDirIfNotExists()

	// If any data available (any file exists), remove all those
	trafficlight.RemoveContents("../data")

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
