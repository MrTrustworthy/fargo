package tests

import (
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/game"
	"testing"
)

func TestIntegration(t *testing.T) {
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
