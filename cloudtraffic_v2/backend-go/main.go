package main
import (
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/websocket"
)

func main() {
	log.Println("Traffic Light App V2 started...")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Printf("Error while serving websocket route ->", err)
	}
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	client := &websocket.Client{
		ID:   timestamp,
		Pool: pool,
		Connection: conn,
		Stopper: nil,
	}
	pool.Register <- client
	websocket.AddClient(timestamp, client)
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
