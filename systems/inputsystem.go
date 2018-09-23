package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/events"
)

type InputSystem struct {
	*ecs.World
	events.MouseTracker
}

func (is *InputSystem) New(world *ecs.World) {
	is.MouseTracker = events.MouseTracker{
		BasicEntity:    ecs.NewBasic(),
		MouseComponent: common.MouseComponent{Track: true},
	}
	AddToMouseSystem(world, &is.MouseTracker)
}

func (is *InputSystem) Update(dt float32) {
	// TODO write a system that takes all events in the mailbox and logs them
	event := events.InputEvent{MouseTracker: is.MouseTracker}
	if engo.Input.Mouse.Action == engo.Press {

		if engo.Input.Mouse.Button == engo.MouseButtonLeft {
			event.Action = events.INPUT_EVENT_ACTION_SELECT
		} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
			event.Action = events.INPUT_EVENT_ACTION_INTERACT
		}

	} else if engo.Input.Button(events.INPUT_EVENT_ACTION_CREATEUNIT).JustPressed() {
		event.Action = events.INPUT_EVENT_ACTION_CREATEUNIT
	} else {
		// don't send out events if nothing has happened
		return
	}
	engo.Mailbox.Dispatch(event)

}

func (is *InputSystem) Remove(e ecs.BasicEntity) {}
