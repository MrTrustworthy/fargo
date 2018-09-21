package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const INPUT_EVENT_NAME = "InputEvent"

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
	message := InputEvent{MouseTracker: is.MouseTracker}
	if engo.Input.Mouse.Action == engo.Press {

		if engo.Input.Mouse.Button == engo.MouseButtonLeft {
			message.Action = "Select"
		} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
			message.Action = "Attack"
		}

	} else if engo.Input.Button("CreateUnit").JustPressed() {
		message.Action = "CreateUnit"
	}
	engo.Mailbox.Dispatch(message)

}

func (is *InputSystem) Remove(e ecs.BasicEntity) {}
