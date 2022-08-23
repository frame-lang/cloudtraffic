package utils

import (
	"log"
)

// All the events received from the UI will be read from here 
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	// Inifite loop to constantly keep an eye on events received from UI for a particular user
	for {
		var UIMsg UIMessage
		err := c.Connection.ReadJSON(&UIMsg)
		if err != nil {
			log.Println("Error while reading messages for connection ID ", c.ID, "Error  -> " ,  err)
			return
		}

		var event string = UIMsg.Event
		log.Println("👉🏻 Event received from the UI 💻 ->", event)
		/*
		 * Events came from UI can be:
		 * 1. createWorkflow
		 * 2. error
		 * 3. restart
		 * 4. end
		 * 5. updateWorkingTimer
		 * 6. updateFlashingTimer
		 */

		if event == "updateWorkingTimer" || event == "updateFlashingTimer" {
			var activeUser *Client = Clients[c.ID]
			if !activeUser.TickInProgress {
				return
			}
			if event == "updateWorkingTimer" {
				activeUser.WorkingTimer = UIMsg.Data
				// Update the exsiting running timer only if state is Working and we updated the Working timer interval
				if UIMsg.State == "working" {
					activeUser.Timer <- true
					activeUser.Timer = setInterval(tick,  activeUser.WorkingTimer, c.ID)
				}
			} else if event == "updateFlashingTimer" {
				activeUser.FlashingTimer = UIMsg.Data
				// Update the exsiting running timer only if state is Error and we updated the Error timer interval
				if UIMsg.State == "error" {
					activeUser.Timer <- true
					activeUser.Timer = setInterval(tick,  activeUser.FlashingTimer, c.ID)
				}
			}
		} else {
			sendEventToTLService(c.ID, event)
		}
	}
}
