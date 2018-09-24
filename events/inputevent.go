package events

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	INPUT_SELECT_EVENT_NAME     = "InputSelectEvent"
	INPUT_INTERACT_EVENT_NAME   = "InputInteractEvent"
	INPUT_CREATEUNIT_EVENT_NAME = "InputCreateunitEvent"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type InputSelectEvent struct {
	MouseTracker
}

func (ie InputSelectEvent) Type() string { return INPUT_SELECT_EVENT_NAME }
func (ie InputSelectEvent) AsLogMessage() string {
	return getInputEventLogMessage(ie.MouseTracker)
}

type InputInteractEvent struct {
	MouseTracker
}

func (ie InputInteractEvent) Type() string { return INPUT_INTERACT_EVENT_NAME }
func (ie InputInteractEvent) AsLogMessage() string {
	return getInputEventLogMessage(ie.MouseTracker)
}

type InputCreateunitEvent struct {
	MouseTracker
}

func (ie InputCreateunitEvent) Type() string { return INPUT_CREATEUNIT_EVENT_NAME }
func (ie InputCreateunitEvent) AsLogMessage() string {
	return getInputEventLogMessage(ie.MouseTracker)
}

func getInputEventLogMessage(tracker MouseTracker) string {
	x, y := PointToXYStrings(engo.Point{tracker.MouseX, tracker.MouseY})
	return " on mouse position (" + x + ":" + y + ")"
}
