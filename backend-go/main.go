package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"

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
		Conn: conn,
		Pool: pool,
	}

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
	
	trafficlight.CreateNewTrafficLight()

	setupRoutes()
	// Create End Point for Web Socket
	// http.HandleFunc("/ws", websocket.WSEndPoint)

	// Create HTTP server
	log.Fatal(http.ListenAndServe(":8000", nil))

	// if err != nil {
	// 	log.Fatal(err)	// }
	// ticker := time.NewTicker(1000 * time.Millisecond)
	//
	// go func() {
	// 	for {
	// 		select {
	// 		case <-stop:
	// 			ticker.Stop()
	// 			mom.Stop()
	// 			finished <- true
	// 			return
	// 		case <-ticker.C:
	// 			fmt.Println("tick")
	// 			mom.Tick()
	// 		}

	// 	}
	// }()

	// stop <- true
	// <-finished
	// fmt.Println("finished")
}
