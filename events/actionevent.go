package events

import (
	"engo.io/engo"
	"strconv"
)

type ActionEvent interface {
	Type() string
	AsLogMessage() string
	GetAction() string
}

func PointToXYStrings(p engo.Point) (x, y string) {
	x = strconv.FormatFloat(float64(p.X), 'f', 3, 64)
	y = strconv.FormatFloat(float64(p.Y), 'f', 3, 64)
	return
}

// this allows for listening to a specific event type only once
// TODO there is absolutely no cleanup, this is horrible
func ListenOnce(messageType, actionType string, handler engo.MessageHandler) {
	alreadyFired := false
	engo.Mailbox.Listen(messageType, func(msg engo.Message) {
		if alreadyFired {
			return
		}
		if amsg, ok := msg.(ActionEvent); !ok || amsg.GetAction() != actionType {
			return
		}

		alreadyFired = true
		handler(msg)
	})
}
