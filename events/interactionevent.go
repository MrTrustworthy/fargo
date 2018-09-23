package events

import (
	"engo.io/engo"
)

const (
	INTERACTION_EVENT_NAME           = "InteractionEvent"
	INTERACTION_EVENT_ACTION_MOVE_TO = "RequestMoveTo"
)

type InteractionEvent struct {
	Target         engo.Point
	Action         string
	StopAtDistance float32
}

func (ae InteractionEvent) Type() string      { return INTERACTION_EVENT_NAME }
func (ae InteractionEvent) GetAction() string { return ae.Action }

func (ae InteractionEvent) AsLogMessage() string {
	x, y := PointToXYStrings(ae.Target)
	return "Action[" + ae.Action + "] on mouse position (" + x + ":" + y
}
