package events

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	INPUT_EVENT_NAME              = "InputEvent"
	INPUT_EVENT_ACTION_SELECT     = "Select"
	INPUT_EVENT_ACTION_INTERACT   = "Interact"
	INPUT_EVENT_ACTION_CREATEUNIT = "CreateUnit"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}
type InputEvent struct {
	MouseTracker
	Action string
}

func (ie InputEvent) Type() string { return INPUT_EVENT_NAME }

func (ie InputEvent) AsLogMessage() string {
	x, y := PointToXYStrings(engo.Point{ie.MouseX, ie.MouseY})
	return "Action[" + ie.Action + "] on mouse position (" + x + ":" + y + ")"
}
