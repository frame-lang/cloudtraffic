package websocket

import (
	"log"
	
    "github.com/frame-lang/frame-demos/persistenttrafficlight/trafficlight"
)

func (c *Client) Read() {
    defer func() {
        c.Pool.Unregister <- c
        c.Conn.Close()
    }()

    for {
        _, p, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        trafficLightMom := trafficlight.TrafficLights[c.ID]

        if string(p) == "start" {
			trafficLightMom.Start()
		} else if string(p) == "error" {
			trafficLightMom.SystemError()
		} else if string(p) == "restart" {
			trafficLightMom.SystemRestart()
		} else if string(p) == "end" {
			trafficLightMom.Stop()
		}
    }
}