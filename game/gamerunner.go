package game

import (
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/scenes"
)

func RunGame() {
	opts := engo.RunOptions{
		Title:          "Hello World",
		Width:          1200,
		Height:         800,
		StandardInputs: true,
	}
	engo.Run(opts, &scenes.WorldScene{})
}
