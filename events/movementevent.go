package events

import (
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	MOVEMENT_REQUESTMOVE_EVENT_NAME = "MovementRequestEvent"
	MOVEMENT_COMPLETED_EVENT_NAME   = "MovementCompletedEvent"
	MOVEMENT_STEP_EVENT_NAME        = "MovementStepEvent"
)

type MovementCompletedEvent struct {
	*entities.Unit
	Successful bool
}

func (me MovementCompletedEvent) Type() string { return MOVEMENT_COMPLETED_EVENT_NAME }

func (me MovementCompletedEvent) AsLogMessage() string {
	x, y := PointToXYStrings(me.Unit.SpaceComponent.Center())
	s := "Aborted"
	if me.Successful {
		s = "Successful"
	}
	return s + " for unit " + me.Unit.Name + " at (" + x + ":" + y + ")"
}

type MovementStepEvent struct {
	*entities.Unit
}

func (me MovementStepEvent) Type() string { return MOVEMENT_STEP_EVENT_NAME }

func (me MovementStepEvent) AsLogMessage() string {
	x, y := PointToXYStrings(me.Unit.SpaceComponent.Center())
	return "for unit " + me.Unit.Name + " at (" + x + ":" + y + ")"
}

type MovementRequestEvent struct {
	Target         engo.Point
	StopAtDistance float32
	Unit           *entities.Unit
}

func (rmte MovementRequestEvent) Type() string { return MOVEMENT_REQUESTMOVE_EVENT_NAME }

func (rmte MovementRequestEvent) AsLogMessage() string {
	x, y := PointToXYStrings(rmte.Target)
	return "on mouse position (" + x + ":" + y
}
