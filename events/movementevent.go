package events

import (
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	MOVEMENT_EVENT_NAME            = "InteractionEvent"
	MOVEMENT_EVENT_ACTION_FINISHED = "MoveCompleted"
	MOVEMENT_EVENT_ACTION_STEP     = "MoveStep"
)

type MovementEvent struct {
	*entities.Unit
	Action string
}

func (ae MovementEvent) Type() string { return MOVEMENT_EVENT_NAME }
