package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/events"
)

const (
	INPUT_CREATE_UNIT_KEY_BIND  = "CreateUnit"
	INPUT_RUN_TESTS_KEY_BIND    = "RunTests"
	INPUT_SHOW_INVENTORY_DIALOG = "ShowInventoryDialog"
	INPUT_HIDE_DIALOG           = "HideDialog"
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
	is.World = world
	is.MouseTracker = MouseTracker{
		BasicEntity:    ecs.NewBasic(),
		MouseComponent: common.MouseComponent{Track: true},
	}
	AddToMouseSystem(world, &is.MouseTracker)
}

func (is *InputSystem) Update(dt float32) {

	selectedUnit := GetCurrentlySelectedUnit(is.World)

	if engo.Input.Mouse.Action == engo.Press {
		if IsDialogUnderMouse(is.World, is.MouseTracker.toPoint()) {
			events.Mailbox.Dispatch(events.DialogClickEvent{Point: is.MouseTracker.toPoint()})
		} else if engo.Input.Mouse.Button == engo.MouseButtonLeft {
			events.Mailbox.Dispatch(events.InputSelectEvent{
				Point: is.MouseTracker.toPoint(),
			})
		} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
			events.Mailbox.Dispatch(events.InputInteractEvent{
				Point: is.MouseTracker.toPoint(),
				Unit: selectedUnit,
			})
		}

	} else if engo.Input.Button(INPUT_CREATE_UNIT_KEY_BIND).JustPressed() {
		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: is.MouseTracker.toPoint()})
	} else if engo.Input.Button(INPUT_RUN_TESTS_KEY_BIND).JustPressed() {
		events.Mailbox.Dispatch(events.TestBasicAttackEvent{})
	} else if engo.Input.Button(INPUT_SHOW_INVENTORY_DIALOG).JustPressed() {
		events.Mailbox.Dispatch(events.ShowInventoryEvent{Unit: selectedUnit})
	} else if engo.Input.Button(INPUT_HIDE_DIALOG).JustPressed() {
		events.Mailbox.Dispatch(events.DialogHideEvent{})
	}

}

func (is *InputSystem) Remove(e ecs.BasicEntity) {}
