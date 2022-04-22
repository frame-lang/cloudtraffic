package websocket

import (
	"log"

	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/trafficlight"
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/cloudpubsub"
	"cloud.google.com/go/pubsub"

)

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	for {
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		trafficLightMom := trafficlight.TrafficLights[c.ID]

		if string(p) == "start" {
			data := pubsub.Message{
				Data: []byte("Data from Go Back-end service"),
				Attributes: map[string]string{
					"clientId": c.ID,
					"event": "start",
				},
			}
			cloudpubsub.Publish("cloud-traffic-347207", "tl-topic", data)
			cloudpubsub.PullMsgs("cloud-traffic-347207", "tl-subscription")
		} else if string(p) == "error" {
			trafficLightMom.SystemError()
		} else if string(p) == "restart" {
			trafficLightMom.SystemRestart()
		} else if string(p) == "end" {
			trafficLightMom.Stop()
		}
	}
}
