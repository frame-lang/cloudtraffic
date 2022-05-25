package main
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/websocket"
)

func main() {
	fmt.Println("Traffic Light App V2 v0.1.0")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	client := &websocket.Client{
		ID:   timestamp,
		Pool: pool,
		Connection: conn,
		Stopper: nil,
	}
	pool.Register <- client
	websocket.AddUser(timestamp, client)
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	go websocket.PullMsgs() 

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}
