package events

import "engo.io/engo"

var ALL_EVENT_NAMES = []string{INPUT_EVENT_NAME, INTERACTION_EVENT_NAME, MOVEMENT_EVENT_NAME, SELECT_EVENT_NAME}

func InitEventLogging(outfunc func(a ...interface{}) (i int, e error)) {

	for _, eventName := range ALL_EVENT_NAMES {
		engo.Mailbox.Listen(eventName, func(msg engo.Message) {
			eventMsg, ok := msg.(ActionEvent)
			if !ok {
				panic("Trying to log an event that isn't an action event, this shouldn't happen!" + msg.Type())
			}
			outfunc(eventMsg.Type(), eventMsg.AsLogMessage())
		})
	}

}
