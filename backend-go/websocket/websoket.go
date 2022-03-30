package websocket

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/frame-lang/frame-demos/persistenttrafficlight/trafficlight"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(conn *websocket.Conn) {
	for {
		trafficlight.SocketConn = conn
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if string(p) == "start" {
			trafficlight.MOM.Start()
		} else if string(p) == "error" {
			trafficlight.MOM.SystemError()
		} else if string(p) == "restart" {
			trafficlight.MOM.SystemRestart()
		} else if string(p) == "end" {
			trafficlight.MOM.Stop()
		}
	}
}

func WSEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Successfully Connected...")
	reader(ws)
}


func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    return conn, nil
}