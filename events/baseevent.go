package events

import (
	"engo.io/engo"
	"strconv"
)

type BaseEvent interface {
	Type() string
	AsLogMessage() string
}

func PointToXYStrings(p engo.Point) (x, y string) {
	x = strconv.FormatFloat(float64(p.X), 'f', 3, 64)
	y = strconv.FormatFloat(float64(p.Y), 'f', 3, 64)
	return
}

// this allows for listening to a specific event type only once
// TODO there is absolutely no cleanup, this is horrible
func ListenOnce(messageType string, handler engo.MessageHandler) {
	alreadyFired := false
	engo.Mailbox.Listen(messageType, func(msg engo.Message) {
		if alreadyFired {
			return
		}
		if _, ok := msg.(BaseEvent); !ok {
			return
		}

		alreadyFired = true
		handler(msg)
	})
}
