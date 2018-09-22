package systems

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

type InputEvent struct {
	MouseTracker
	Action string
}

func (ie InputEvent) Type() string { return INPUT_EVENT_NAME }

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type InputSystem struct {
	*ecs.World
	MouseTracker
}

func (is *InputSystem) New(world *ecs.World) {
	is.MouseTracker = MouseTracker{
		BasicEntity:    ecs.NewBasic(),
		MouseComponent: common.MouseComponent{Track: true},
	}
	AddToMouseSystem(world, &is.MouseTracker)
}

func (is *InputSystem) Update(dt float32) {
	// TODO write a system that takes all events in the mailbox and logs them
	event := InputEvent{MouseTracker: is.MouseTracker}
	if engo.Input.Mouse.Action == engo.Press {

		if engo.Input.Mouse.Button == engo.MouseButtonLeft {
			event.Action = INPUT_EVENT_ACTION_SELECT
		} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
			event.Action = INPUT_EVENT_ACTION_INTERACT
		}

	} else if engo.Input.Button(INPUT_EVENT_ACTION_CREATEUNIT).JustPressed() {
		event.Action = INPUT_EVENT_ACTION_CREATEUNIT
	}
	engo.Mailbox.Dispatch(event)

}

func (is *InputSystem) Remove(e ecs.BasicEntity) {}
