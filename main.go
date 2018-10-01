package main

import (
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/game"
)

func main() {
	go logMsg(events.CaptureChannel)
	game.RunGame()
}

func logMsg(eventChan chan events.BaseEvent) {
	for true {
		msg := <-eventChan
		if msg.Type() != events.MOVEMENT_STEP_EVENT_NAME {
			fmt.Println(msg.Type(), msg.AsLogMessage())
		}
	}
}
