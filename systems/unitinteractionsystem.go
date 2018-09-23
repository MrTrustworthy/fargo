package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
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
				return
			} else {

				dispatchMoveTo(unit.Center().X, unit.Center().Y, 250)

			}
		} else {
			dispatchMoveTo(imsg.MouseX, imsg.MouseY, 0)
		}

	}
}

func dispatchMoveTo(x, y, dist float32) {
	engo.Mailbox.Dispatch(events.InteractionEvent{
		Target:         engo.Point{X: x, Y: y},
		Action:         events.INTERACTION_EVENT_ACTION_MOVE_TO,
		StopAtDistance: dist,
	})
}

func (uis *UnitInteractionSystem) Update(dt float32) {}

func (uis *UnitInteractionSystem) Remove(e ecs.BasicEntity) {}
