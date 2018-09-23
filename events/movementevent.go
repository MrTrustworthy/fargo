package events

import (
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	MOVEMENT_EVENT_NAME            = "MovementEvent"
	MOVEMENT_EVENT_ACTION_FINISHED = "MoveCompleted"
	MOVEMENT_EVENT_ACTION_STEP     = "MoveStep"
)

type MovementEvent struct {
	*entities.Unit
	Action string
}

func (me MovementEvent) Type() string { return MOVEMENT_EVENT_NAME }

func (me MovementEvent) AsLogMessage() string {
	x, y := PointToXYStrings(me.Unit.SpaceComponent.Position)
	return "Action[" + me.Action + "] for unit " + me.Unit.Name + " at (" + x + ":" + y + ")"
}
