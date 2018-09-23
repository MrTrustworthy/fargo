package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitInteractionSystem struct {
	*ecs.World
}

func (uis *UnitInteractionSystem) New(world *ecs.World) {
	uis.World = world
	engo.Mailbox.Listen(events.INPUT_EVENT_NAME, uis.getHandleInputEvent())
}

func (uis *UnitInteractionSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputEvent)
		if !ok || imsg.Action != events.INPUT_EVENT_ACTION_INTERACT {
			return
		}

		if unit, err := FindUnitUnderMouse(uis.World, &imsg.MouseTracker); err == nil {
			if unit == GetCurrentlySelectedUnit(uis.World) {
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
	engo.Mailbox.Dispatch(events.InteractionEvent{
		Target: engo.Point{X: x, Y: y},
		Action: events.INTERACTION_EVENT_ACTION_MOVE,
	})
}

func (uis *UnitInteractionSystem) Update(dt float32) {}

func (uis *UnitInteractionSystem) Remove(e ecs.BasicEntity) {}
