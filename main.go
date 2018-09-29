package main

import (
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/game"
)

func main() {
	eventChan := make(chan events.BaseEvent)
	go logMsg(eventChan)
	game.RunGame(eventChan)
}

func logMsg(eventChan chan events.BaseEvent) {
	for true {
		msg := <-eventChan
		fmt.Println(msg.Type(), msg.AsLogMessage())
	}
}
