package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/events"
)

const (
	INPUT_CREATE_UNIT_KEY_BIND = "CreateUnit"
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
	if engo.Input.Mouse.Action == engo.Press {

		if engo.Input.Mouse.Button == engo.MouseButtonLeft {
			events.Mailbox.Dispatch(events.InputSelectEvent{
				MouseTracker: is.MouseTracker,
			})
		} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
			events.Mailbox.Dispatch(events.InputInteractEvent{
				MouseTracker: is.MouseTracker,
			})
		}

	} else if engo.Input.Button(INPUT_CREATE_UNIT_KEY_BIND).JustPressed() {
		events.Mailbox.Dispatch(events.InputCreateunitEvent{
			MouseTracker: is.MouseTracker,
		})
	}

}

func (is *InputSystem) Remove(e ecs.BasicEntity) {}
