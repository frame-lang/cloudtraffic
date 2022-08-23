package main
import (
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/frame-lang/utils-service/utils"
)

func main() {
	log.Println("🚦 Traffic Light App V3 starting...")
	setupRoutes()

	rabbitConnection := utils.ConnectToRabbitMQ();
	defer rabbitConnection.Conn.Close()
	defer rabbitConnection.Channel.Close()

	go utils.CreateUIResponseQueueSubscriber(rabbitConnection.Channel)
	go utils.CreateTimerQueueSubscriber(rabbitConnection.Channel)

	log.Fatal(http.ListenAndServe(":9070", nil))
}

func serveWs(pool *utils.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket Endpoint Hit")
	conn, err := utils.Upgrade(w, r)
	if err != nil {
		log.Printf("Error while serving websocket route ->", err)
	}
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	client := &utils.Client{
		ID:				timestamp,
		Pool:			pool,
		Connection: 	conn,
		Timer:			nil,
		TickInProgress: false,
		WorkingTimer:	"2s",
		FlashingTimer:	"1s",
	}
	pool.Register <- client
	utils.AddClient(timestamp, client)
	client.Read()
}

func setupRoutes() {
	pool := utils.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}
