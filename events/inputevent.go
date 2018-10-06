package events

import (
	"engo.io/engo"
)

const (
	INPUT_SELECT_EVENT_NAME     = "InputSelectEvent"
	INPUT_INTERACT_EVENT_NAME   = "InputInteractEvent"
	INPUT_CREATEUNIT_EVENT_NAME = "InputCreateunitEvent"
)

// TODO remove mouse tracker, replace with point
type InputSelectEvent struct {
	engo.Point
}

func (ie InputSelectEvent) Type() string { return INPUT_SELECT_EVENT_NAME }
func (ie InputSelectEvent) AsLogMessage() string {
	return getInputEventLogMessage(ie.Point)
}

// TODO remove mouse tracker, replace with point
type InputInteractEvent struct {
	engo.Point
}

func (ie InputInteractEvent) Type() string { return INPUT_INTERACT_EVENT_NAME }
func (ie InputInteractEvent) AsLogMessage() string {
	return getInputEventLogMessage(ie.Point)
}

type InputCreateunitEvent struct {
	engo.Point
}

func (ie InputCreateunitEvent) Type() string { return INPUT_CREATEUNIT_EVENT_NAME }
func (ie InputCreateunitEvent) AsLogMessage() string {
	return getInputEventLogMessage(ie.Point)
}

func getInputEventLogMessage(p engo.Point) string {
	x, y := PointToXYStrings(p)
	return " on mouse position (" + x + ":" + y + ")"
}
