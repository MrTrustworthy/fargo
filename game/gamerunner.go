package game

import (
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/scenes"
)

func RunGame(eventChannel chan<- events.BaseEvent) {
	opts := engo.RunOptions{
		Title:          "Hello World",
		Width:          1200,
		Height:         800,
		StandardInputs: true,
	}
	engo.Run(opts, &scenes.WorldScene{
		EventChannel: eventChannel,
	})
}
