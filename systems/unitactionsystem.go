package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
)

const ACTION_EVENT_NAME = "ActionEvent"

type ActionEvent struct {
	Target engo.Point
	Action string
}

func (ae ActionEvent) Type() string { return ACTION_EVENT_NAME }

type UnitActionSystem struct {
	*ecs.World
}

func (uas *UnitActionSystem) New(world *ecs.World) {
	uas.World = world
	engo.Mailbox.Listen(INPUT_EVENT_NAME, uas.getHandleInputEvent())
}

func (uas *UnitActionSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(InputEvent)
		if !ok || imsg.Action != "Interact" {
			return
		}

		if unit, err := FindUnitUnderMouse(uas.World, &imsg.MouseTracker); err == nil {
			if unit == GetCurrentlySelectedUnit(uas.World) {
				fmt.Println("TODO: Handle action on selected unit itself")
			} else {
				fmt.Println("TODO: Handle action on other unit")
			}
		} else {
			dispatchMoveTo(imsg.MouseX, imsg.MouseY)
		}

	}
}

func dispatchMoveTo(x, y float32) {
	engo.Mailbox.Dispatch(ActionEvent{
		Target: engo.Point{X: x, Y: y},
		Action: "Move",
	})
}

func (uas *UnitActionSystem) Update(dt float32) {}

func (uas *UnitActionSystem) Remove(e ecs.BasicEntity) {}
