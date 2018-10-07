package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/events"
)

const (
	INPUT_CREATE_UNIT_KEY_BIND = "CreateUnit"
	INPUT_RUN_TESTS_KEY_BIND   = "RunTests"
	INPUT_SHOW_DIALOG          = "ShowDialog"
	INPUT_HIDE_DIALOG          = "HideDialog"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

func (mt *MouseTracker) toPoint() engo.Point {
	return engo.Point{X: mt.MouseX, Y: mt.MouseY}
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
	if engo.Input.Mouse.Action == engo.Press {

		if engo.Input.Mouse.Button == engo.MouseButtonLeft {
			events.Mailbox.Dispatch(events.InputSelectEvent{
				Point: is.MouseTracker.toPoint(),
			})
		} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
			events.Mailbox.Dispatch(events.InputInteractEvent{
				Point: is.MouseTracker.toPoint(),
			})
		}

	} else if engo.Input.Button(INPUT_CREATE_UNIT_KEY_BIND).JustPressed() {
		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: is.MouseTracker.toPoint()})
	} else if engo.Input.Button(INPUT_RUN_TESTS_KEY_BIND).JustPressed() {
		events.Mailbox.Dispatch(events.TestBasicAttackEvent{})
	} else if engo.Input.Button(INPUT_SHOW_DIALOG).JustPressed() {
		events.Mailbox.Dispatch(events.DialogShowEvent{})
	}else if engo.Input.Button(INPUT_HIDE_DIALOG).JustPressed() {
		events.Mailbox.Dispatch(events.DialogHideEvent{})
	}

}

func (is *InputSystem) Remove(e ecs.BasicEntity) {}
