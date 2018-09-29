package events

import "engo.io/engo"

var ALL_EVENT_NAMES = []string{
	COLLISON_EVENT_NAME,
	MOVEMENT_REQUESTMOVE_EVENT_NAME,
	MOVEMENT_COMPLETED_EVENT_NAME,
	//MOVEMENT_STEP_EVENT_NAME, too many...
	ABILITY_REQUESTUSE_EVENT_NAME,
	SELECTION_SELECTED_EVENT_NAME,
	SELECTION_DESELECTED_EVENT_NAME,
	INPUT_SELECT_EVENT_NAME,
	INPUT_INTERACT_EVENT_NAME,
	INPUT_CREATEUNIT_EVENT_NAME,
	UNIT_REQUEST_DAMAGE_EVENT,
	UNIT_DEATH_EVENT,
	ABILITY_REQUESTUSE_EVENT_NAME,
	ABILITY_COMPLETED_EVENT_NAME,
	ABILITY_ABORT_EVENT_NAME,
	UNIT_ATTRIBUTE_CHANGE_EVENT,
}

func InitEventCapturing(channel chan<- BaseEvent) {

	for _, eventName := range ALL_EVENT_NAMES {
		engo.Mailbox.Listen(eventName, func(msg engo.Message) {
			eventMsg, ok := msg.(BaseEvent)
			if !ok {
				panic("Trying to log an event that isn't a BaseEvent, this shouldn't happen!" + msg.Type())
			}
			channel <- eventMsg
		})
	}

}
